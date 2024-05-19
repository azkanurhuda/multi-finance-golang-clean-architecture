package merchant

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

type MerchantUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (m MerchantUseCase) Create(ctx context.Context, request *model.CreateMerchantRequest) (*model.MerchantResponse, error) {
	tx := m.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := m.Validate.Struct(request)
	if err != nil {
		m.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	merchant := &entity.Merchant{
		ID:        uuid.NewString(),
		Name:      request.Name,
		CreatedAt: time.Now().Local(),
		UpdatedAt: time.Now().Local(),
	}

	if err := m.Repository.Merchant.Create(tx, merchant); err != nil {
		m.Log.Warnf("Failed create mercahnt to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		m.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.MerchantToResponse(merchant), nil
}

func (m MerchantUseCase) GetByID(ctx context.Context, id string) (*model.MerchantResponse, error) {
	tx := m.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	merchant, err := m.Repository.Merchant.GetByID(tx, id)
	if err != nil {
		m.Log.Warnf("Failed get merchant by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		m.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.MerchantToResponse(merchant), nil
}

func (m MerchantUseCase) List(ctx context.Context) ([]model.MerchantResponse, error) {
	tx := m.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	merchant, err := m.Repository.Merchant.GetList(tx)
	if err != nil {
		m.Log.Warnf("Failed get merchant list in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		m.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ListMerchantToResponse(merchant), nil
}

func (m MerchantUseCase) Delete(ctx context.Context, id string) (bool, error) {
	tx := m.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := m.Repository.Merchant.DeleteByID(tx, id)
	if err != nil {
		m.Log.Warnf("Failed delete merchant in database : %+v", err)
		return false, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		m.Log.Warnf("Failed commit transaction : %+v", err)
		return false, fiber.ErrInternalServerError
	}

	return true, nil
}

func NewMerchantUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &MerchantUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
