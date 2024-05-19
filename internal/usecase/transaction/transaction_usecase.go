package transaction

import (
	"context"
	"fmt"
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

type TransactionUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (t TransactionUseCase) CreateByAdmin(ctx context.Context, request *model.CreateTransactionByAdminRequest) (*model.TransactionResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := t.Validate.Struct(request)
	if err != nil {
		t.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user, err := t.Repository.User.FindByID(tx, request.UserID)
	if err != nil {
		t.Log.Warnf("asset not found : %+v", err)
		return nil, fiber.ErrNotFound
	}

	assetData, err := t.Repository.Asset.GetByID(tx, request.AssetID)
	if err != nil {
		t.Log.Warnf("asset not found : %+v", err)
		return nil, fiber.ErrNotFound
	}

	otr := assetData.OTR
	interestAmount := assetData.OTR * 0.08
	adminFee := assetData.AdminFee
	totalPayment := otr + interestAmount + adminFee

	if request.PaymentMethod == "kredit" {
		transaction := &entity.Transaction{
			ID:             uuid.NewString(),
			UserID:         request.UserID,
			AssetID:        request.AssetID,
			ContractNumber: request.ContractNumber,
			TotalPayment:   totalPayment,
			PaymentMethod:  request.PaymentMethod,
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			CreditID:       request.CreditID,
			User:           entity.User{},
			Asset:          entity.Asset{},
			Credit:         nil,
		}

		if err := t.Repository.Transaction.Create(tx, transaction); err != nil {
			t.Log.Warnf("Failed create transaction to database : %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		response, err := t.Repository.Transaction.FindByID(tx, transaction.ID)
		if err != nil {
			t.Log.Warnf("error get detail transaction : %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		if err := tx.Commit().Error; err != nil {
			t.Log.Warnf("Failed commit transaction : %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		return converter.TransactionToResponse(response), nil
	}

	var transaction *entity.Transaction
	if request.CreditID != nil {
		transaction = &entity.Transaction{
			ID:             uuid.NewString(),
			UserID:         user.ID,
			AssetID:        request.AssetID,
			ContractNumber: request.ContractNumber,
			TotalPayment:   totalPayment,
			PaymentMethod:  request.PaymentMethod,
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			CreditID:       request.CreditID,
		}
	}

	transaction = &entity.Transaction{
		ID:             uuid.NewString(),
		UserID:         user.ID,
		AssetID:        request.AssetID,
		ContractNumber: request.ContractNumber,
		TotalPayment:   totalPayment,
		PaymentMethod:  request.PaymentMethod,
		CreatedAt:      time.Now().Local(),
		UpdatedAt:      time.Now().Local(),
	}

	if err := t.Repository.Transaction.Create(tx, transaction); err != nil {
		t.Log.Warnf("Failed create transaction to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	fmt.Println("111111")
	response, err := t.Repository.Transaction.FindByID(tx, transaction.ID)
	if err != nil {
		t.Log.Warnf("error get detail transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}
	fmt.Println("------")

	if err := tx.Commit().Error; err != nil {
		t.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TransactionToResponse(response), nil
}

func (t TransactionUseCase) Create(ctx context.Context, request *model.CreateTransactionRequest, auth model.Auth) (*model.TransactionResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := t.Validate.Struct(request)
	if err != nil {
		t.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	_, err = t.Repository.User.FindByID(tx, auth.UserID)
	if err != nil {
		t.Log.Warnf("asset not found : %+v", err)
		return nil, fiber.ErrNotFound
	}

	assetData, err := t.Repository.Asset.GetByID(tx, request.AssetID)
	if err != nil {
		t.Log.Warnf("asset not found : %+v", err)
		return nil, fiber.ErrNotFound
	}

	otr := assetData.OTR
	interestAmount := assetData.OTR * 0.08
	adminFee := assetData.AdminFee
	totalPayment := otr + interestAmount + adminFee

	if request.PaymentMethod == "kredit" {
		transaction := &entity.Transaction{
			ID:             uuid.NewString(),
			UserID:         auth.UserID,
			AssetID:        request.AssetID,
			ContractNumber: request.ContractNumber,
			TotalPayment:   totalPayment,
			PaymentMethod:  request.PaymentMethod,
			CreatedAt:      time.Now().Local(),
			UpdatedAt:      time.Now().Local(),
			CreditID:       request.CreditID,
			User:           entity.User{},
			Asset:          entity.Asset{},
			Credit:         nil,
		}

		if err := t.Repository.Transaction.Create(tx, transaction); err != nil {
			t.Log.Warnf("Failed create transaction to database : %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		response, err := t.Repository.Transaction.FindByID(tx, transaction.ID)
		if err != nil {
			t.Log.Warnf("error get detail transaction : %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		if err := tx.Commit().Error; err != nil {
			t.Log.Warnf("Failed commit transaction : %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		return converter.TransactionToResponse(response), nil
	}

	transaction := &entity.Transaction{
		ID:             uuid.NewString(),
		UserID:         auth.UserID,
		AssetID:        request.AssetID,
		ContractNumber: request.ContractNumber,
		TotalPayment:   totalPayment,
		PaymentMethod:  request.PaymentMethod,
		CreatedAt:      time.Now().Local(),
		UpdatedAt:      time.Now().Local(),
		CreditID:       request.CreditID,
		User:           entity.User{},
		Asset:          entity.Asset{},
		Credit:         nil,
	}

	if err := t.Repository.Transaction.Create(tx, transaction); err != nil {
		t.Log.Warnf("Failed create transaction to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	response, err := t.Repository.Transaction.FindByID(tx, transaction.ID)
	if err != nil {
		t.Log.Warnf("error get detail transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		t.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TransactionToResponse(response), nil
}

func NewTransactionUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &TransactionUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
