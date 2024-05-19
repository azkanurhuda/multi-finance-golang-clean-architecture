package credit_payment

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

type CreditPaymentUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (c CreditPaymentUseCase) Create(ctx context.Context, request *model.CreateCreditPaymentRequest) (*model.CreditPaymentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	credit, err := c.Repository.Credit.GetByID(tx, request.CreditID)
	if err != nil {
		c.Log.Warnf("Failed get credit by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	listCredit, err := c.Repository.CreditPayment.GetListByCreditID(tx, request.CreditID)
	if err != nil {
		c.Log.Warnf("Failed get credit by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	totalPay := credit.MonthlyInstallment * float64(credit.Tenor)
	var donePayment float64
	for _, v := range listCredit {
		donePayment += v.PaymentAmount
	}

	if donePayment > totalPay {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "anda sudah membayar semua kredit")
	}

	if request.PaymentAmount != credit.MonthlyInstallment {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "anda harus bayar sesuai tagihan")
	}

	if len(listCredit) == credit.Tenor {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "anda sudah tidak punya tagihan")
	}

	data := &entity.CreditPayment{
		ID:            uuid.NewString(),
		CreditID:      request.CreditID,
		PaymentAmount: request.PaymentAmount,
		CreatedAt:     time.Now().Local(),
		UpdatedAt:     time.Now().Local(),
	}

	if err := c.Repository.CreditPayment.Create(tx, data); err != nil {
		c.Log.Warnf("Failed create consumer to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	res, err := c.Repository.CreditPayment.GetByID(tx, data.ID)
	if err != nil {
		c.Log.Warnf("Failed get credit payment by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CreditPaymentToResponse(res), nil
}

func (c CreditPaymentUseCase) GetByID(ctx context.Context, id string) (*model.CreditPaymentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	creditPayment, err := c.Repository.CreditPayment.GetByID(tx, id)
	if err != nil {
		c.Log.Warnf("Failed get credit payment by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	_, err = c.Repository.Credit.GetByID(tx, creditPayment.CreditID)
	if err != nil {
		c.Log.Warnf("Failed get credit payment by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	res, err := c.Repository.CreditPayment.GetByID(tx, id)
	if err != nil {
		c.Log.Warnf("Failed get credit payment by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.CreditPaymentToResponse(res), nil
}

func (c CreditPaymentUseCase) List(ctx context.Context) ([]model.CreditPaymentResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	credit, err := c.Repository.CreditPayment.GetList(tx)
	if err != nil {
		c.Log.Warnf("Failed get credit payment list in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ListCreditPaymentToResponse(credit), nil
}

func (c CreditPaymentUseCase) Delete(ctx context.Context, id string) (bool, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Repository.CreditPayment.DeleteByID(tx, id)
	if err != nil {
		c.Log.Warnf("Failed delete credit payment in database : %+v", err)
		return false, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return false, fiber.ErrInternalServerError
	}

	return true, nil
}

func NewCreditPaymentUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &CreditPaymentUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
