package transaction

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
	"strconv"
	"time"
)

type TransactionUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (t TransactionUseCase) Create(ctx context.Context, request *model.CreateTransactionRequest, nik string) (*model.TransactionResponse, error) {
	tx := t.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
	adminFee := 10000

	err := t.Validate.Struct(request)
	if err != nil {
		t.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	data, _ := t.Repository.ConsumerLimit.FindByNIKAndTenor(tx, nik, strconv.FormatInt(request.InstallmentAmount, 10))
	if data == nil {
		t.Log.Warnf("can not find consumer limit by nik and tenor : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	otrAdmin := request.OTR + int64(adminFee)
	if otrAdmin > data.Limits {
		t.Log.Warnf("total OTR + biaya admin lebih dari limit tenornya")
		return nil, fiber.ErrBadRequest
	}

	bunga := float64(otrAdmin) * 0.02

	transaction := &entity.Transaction{
		ID:                uuid.NewString(),
		ContractNumber:    "TRX" + uuid.NewString(),
		NIK:               nik,
		OTR:               request.OTR,
		AdminFee:          int64(adminFee),
		InstallmentAmount: request.InstallmentAmount,
		AmountOfInterest:  int64(bunga),
		AssetName:         request.AssetName,
		CreatedAt:         time.Now().Local(),
		UpdatedAt:         time.Now().Local(),
	}

	if err := t.Repository.Transaction.Create(tx, transaction); err != nil {
		t.Log.Warnf("Failed create consumer to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		t.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.TransactionToResponse(transaction), nil
}

func NewTransactionUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &TransactionUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
