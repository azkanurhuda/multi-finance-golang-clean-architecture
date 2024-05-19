package credit_limit

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

type CreditLimitUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (c CreditLimitUseCase) Create(ctx context.Context, request *model.CreateCreditLimitRequest) (*model.CreditLimitResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user, err := c.Repository.User.FindByID(tx, request.UserID)
	if err != nil {
		c.Log.Warnf("Failed find user to database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	data := &entity.CreditLimit{
		ID:          uuid.NewString(),
		UserID:      request.UserID,
		Tenor:       request.Tenor,
		CreditLimit: request.Limit,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   time.Now().Local(),
	}

	if err := c.Repository.CreditLimit.Create(tx, data); err != nil {
		c.Log.Warnf("Failed create credit limit to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CreditLimitToResponse(data, user), nil
}

func (c CreditLimitUseCase) GetByID(ctx context.Context, id string) (*model.CreditLimitResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	creditLimit, err := c.Repository.CreditLimit.GetByID(tx, id)
	if err != nil {
		c.Log.Warnf("Failed get asset by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	user, err := c.Repository.User.FindByID(tx, creditLimit.UserID)
	if err != nil {
		c.Log.Warnf("Failed get asset by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CreditLimitToResponse(creditLimit, user), nil
}

func (c CreditLimitUseCase) List(ctx context.Context) ([]model.CreditLimitResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	creditLimit, err := c.Repository.CreditLimit.GetList(tx)
	if err != nil {
		c.Log.Warnf("Failed get credit limit list in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ListCreditLimitToResponse(creditLimit), nil
}

func (c CreditLimitUseCase) Delete(ctx context.Context, id string) (bool, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Repository.CreditLimit.DeleteByID(tx, id)
	if err != nil {
		c.Log.Warnf("Failed delete credit limit in database : %+v", err)
		return false, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return false, fiber.ErrInternalServerError
	}

	return true, nil
}

func NewCreditLimitUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &CreditLimitUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
