package credit

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model/converter"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"time"
)

type CreditUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (c CreditUseCase) Create(ctx context.Context, request *model.CreateCreditRequest, auth model.Auth) (*model.CreditResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	transaction, err := c.Repository.Transaction.GetByID(tx, request.TransactionID)
	if err != nil {
		c.Log.Warnf("Failed get transaction by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	var kredit float64
	// Check Credit Limit
	creditLimit, _ := c.Repository.CreditLimit.GetByUserIDAndTenor(tx, auth.UserID, request.Tenor)
	if creditLimit == nil {
		kredit = 0
	}

	assetData, err := c.Repository.Asset.GetByID(tx, transaction.AssetID)
	if err != nil {
		c.Log.Warnf("asset not found : %+v", err)
		return nil, fiber.ErrNotFound
	}

	totalKredit := assetData.OTR - kredit
	interestAmount := assetData.OTR * 0.08
	totalKredit += interestAmount
	monthlyInstallment := totalKredit / float64(request.Tenor)

	credit := &entity.Credit{
		ID:                 uuid.NewString(),
		TransactionID:      request.TransactionID,
		Tenor:              request.Tenor,
		CreditLimit:        creditLimit.CreditLimit,
		MonthlyInstallment: monthlyInstallment,
		InterestAmount:     interestAmount,
		CreatedAt:          time.Now().Local(),
		UpdatedAt:          time.Now().Local(),
	}

	if err := c.Repository.Credit.Create(tx, credit); err != nil {
		c.Log.Warnf("Failed create credit to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	transaction.CreditID = &credit.ID
	if err := c.Repository.Transaction.UpdateCreditID(tx, transaction); err != nil {
		c.Log.Warnf("Failed create credit to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	res, err := c.Repository.Credit.GetByID(tx, credit.ID)
	if err != nil {
		c.Log.Warnf("Failed get credit to database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CreditToResponse(res), nil
}

func (c CreditUseCase) GetByID(ctx context.Context, id string) (*model.CreditResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	credit, err := c.Repository.Credit.GetByID(tx, id)
	if err != nil {
		c.Log.Warnf("Failed get credit by id to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	_, err = c.Repository.Transaction.GetByID(tx, credit.TransactionID)
	if err != nil {
		c.Log.Warnf("Failed get transaction by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	res, err := c.Repository.Credit.GetByID(tx, credit.ID)
	if err != nil {
		c.Log.Warnf("Failed get credit to database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CreditToResponse(res), nil
}

func (c CreditUseCase) List(ctx context.Context) ([]model.CreditResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	credit, err := c.Repository.Credit.GetList(tx)
	if err != nil {
		c.Log.Warnf("Failed get credit list in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ListCreditToResponse(credit), nil
}

func (c CreditUseCase) Delete(ctx context.Context, id string) (bool, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Repository.Credit.DeleteByID(tx, id)
	if err != nil {
		c.Log.Warnf("Failed delete credit in database : %+v", err)
		return false, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return false, fiber.ErrInternalServerError
	}

	return true, nil
}

func NewCreditUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &CreditUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
