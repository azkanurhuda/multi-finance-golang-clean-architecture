package asset

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

type AssetUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (a AssetUseCase) Create(ctx context.Context, request *model.CreateAssetRequest) (*model.AssetResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := a.Validate.Struct(request)
	if err != nil {
		a.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	merchant, err := a.Repository.Merchant.GetByID(tx, request.MerchantID)
	if err != nil {
		a.Log.Warnf("Failed get merchant by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	asset := &entity.Asset{
		ID:          uuid.NewString(),
		MerchantID:  request.MerchantID,
		Name:        request.Name,
		OTR:         request.OTR,
		AdminFee:    request.AdminFee,
		Description: request.Description,
		CreatedAt:   time.Now().Local(),
		UpdatedAt:   time.Now().Local(),
	}

	if err := a.Repository.Asset.Create(tx, asset); err != nil {
		a.Log.Warnf("Failed create asset to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.AssetToResponse(asset, merchant), nil
}

func (a AssetUseCase) GetByID(ctx context.Context, id string) (*model.AssetResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	asset, err := a.Repository.Asset.GetByID(tx, id)
	if err != nil {
		a.Log.Warnf("Failed get asset by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	merchant, err := a.Repository.Merchant.GetByID(tx, asset.MerchantID)
	if err != nil {
		a.Log.Warnf("Failed get asset by id in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.AssetToResponse(asset, merchant), nil
}

func (a AssetUseCase) List(ctx context.Context) ([]model.AssetResponse, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	asset, err := a.Repository.Asset.GetList(tx)
	if err != nil {
		a.Log.Warnf("Failed get asset list in database : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ListAssetToResponse(asset), nil
}

func (a AssetUseCase) Delete(ctx context.Context, id string) (bool, error) {
	tx := a.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := a.Repository.Asset.DeleteByID(tx, id)
	if err != nil {
		a.Log.Warnf("Failed delete asset in database : %+v", err)
		return false, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		a.Log.Warnf("Failed commit transaction : %+v", err)
		return false, fiber.ErrInternalServerError
	}

	return true, nil
}

func NewAssetUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &AssetUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
