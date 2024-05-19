package consumer

import (
	"context"
	"encoding/base64"
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

type ConsumerUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func (c ConsumerUseCase) CreateByAdmin(ctx context.Context, request *model.CreateConsumerByAdminRequest) (*model.ConsumerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	_, err = c.Repository.User.FindByID(tx, request.UserID)
	if err != nil {
		c.Log.Warnf("User not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	data, err := c.Repository.Consumer.FindByUserID(tx, request.UserID)
	if data != nil {
		c.Log.Warnf("User already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	const dateFormat = "02-01-2006"
	dateOfBirth, err := time.Parse(dateFormat, request.DateOfBirth)
	if err != nil {
		c.Log.Warnf("Failed to parse date of birth: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	idCardPhoto, err := base64.StdEncoding.DecodeString(request.IDCardPhoto)
	if err != nil {
		c.Log.Warnf("Failed to decode id card photo: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	selfiePhoto, err := base64.StdEncoding.DecodeString(request.SelfiePhoto)
	if err != nil {
		c.Log.Warnf("Failed to decode selfie photo: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Convert byte slices to base64-encoded strings
	idCardPhotoString := base64.StdEncoding.EncodeToString(idCardPhoto)
	selfiePhotoString := base64.StdEncoding.EncodeToString(selfiePhoto)

	consumer := &entity.Consumer{
		ID:           uuid.NewString(),
		UserID:       request.UserID,
		NIK:          request.NIK,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		PhoneNumber:  request.PhoneNumber,
		Address:      request.Address,
		PlaceOfBirth: request.PlaceOfBirth,
		DateOfBirth:  dateOfBirth,
		Salary:       request.Salary,
		IDCardPhoto:  idCardPhotoString,
		SelfiePhoto:  selfiePhotoString,
		CreatedAt:    time.Now().Local(),
		UpdatedAt:    time.Now().Local(),
		User:         entity.User{},
	}

	if err := c.Repository.Consumer.Create(tx, consumer); err != nil {
		c.Log.Warnf("Failed create consumer to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Add Credit Limit For New User
	var firstLimit float64
	firstLimit = 4000000
	for i := 1; i <= 4; i++ {
		creditLimit := &entity.CreditLimit{
			ID:          uuid.NewString(),
			UserID:      request.UserID,
			Tenor:       i,
			CreditLimit: firstLimit,
			CreatedAt:   time.Now().Local(),
			UpdatedAt:   time.Now().Local(),
		}

		if err := c.Repository.CreditLimit.Create(tx, creditLimit); err != nil {
			c.Log.Warnf("Failed create credit limit to database : %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		firstLimit += 2000000
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user, err := c.Repository.User.FindByID(c.DB.WithContext(ctx), request.UserID)
	if data != nil {
		c.Log.Warnf("User not found exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	user.Token = ""

	return converter.ConsumerToResponse(consumer, user), nil
}

func (c ConsumerUseCase) Create(ctx context.Context, request *model.CreateConsumerRequest, auth model.Auth) (*model.ConsumerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	data, err := c.Repository.Consumer.FindByUserID(tx, auth.UserID)
	if data != nil {
		c.Log.Warnf("User already exists : %+v", err)
		return nil, fiber.ErrConflict
	}

	const dateFormat = "02-01-2006"
	dateOfBirth, err := time.Parse(dateFormat, request.DateOfBirth)
	if err != nil {
		c.Log.Warnf("Failed to parse date of birth: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	idCardPhoto, err := base64.StdEncoding.DecodeString(request.IDCardPhoto)
	if err != nil {
		c.Log.Warnf("Failed to decode id card photo: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	selfiePhoto, err := base64.StdEncoding.DecodeString(request.SelfiePhoto)
	if err != nil {
		c.Log.Warnf("Failed to decode selfie photo: %+v", err)
		return nil, fiber.ErrBadRequest
	}

	// Convert byte slices to base64-encoded strings
	idCardPhotoString := base64.StdEncoding.EncodeToString(idCardPhoto)
	selfiePhotoString := base64.StdEncoding.EncodeToString(selfiePhoto)

	consumer := &entity.Consumer{
		ID:           uuid.NewString(),
		UserID:       auth.UserID,
		NIK:          request.NIK,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		PhoneNumber:  request.PhoneNumber,
		Address:      request.Address,
		PlaceOfBirth: request.PlaceOfBirth,
		DateOfBirth:  dateOfBirth,
		Salary:       request.Salary,
		IDCardPhoto:  idCardPhotoString,
		SelfiePhoto:  selfiePhotoString,
		CreatedAt:    time.Now().Local(),
		UpdatedAt:    time.Now().Local(),
		User:         entity.User{},
	}

	if err := c.Repository.Consumer.Create(tx, consumer); err != nil {
		c.Log.Warnf("Failed create consumer to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	// Add Credit Limit For New User
	var firstLimit float64
	firstLimit = 4000000
	for i := 1; i <= 4; i++ {
		creditLimit := &entity.CreditLimit{
			ID:          uuid.NewString(),
			UserID:      auth.UserID,
			Tenor:       i,
			CreditLimit: firstLimit,
			CreatedAt:   time.Now().Local(),
			UpdatedAt:   time.Now().Local(),
		}

		if err := c.Repository.CreditLimit.Create(tx, creditLimit); err != nil {
			c.Log.Warnf("Failed create credit limit to database : %+v", err)
			return nil, fiber.ErrInternalServerError
		}

		firstLimit += 2000000
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user, err := c.Repository.User.FindByID(c.DB.WithContext(ctx), auth.UserID)
	if data != nil {
		c.Log.Warnf("User not found exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	user.Token = ""

	return converter.ConsumerToResponse(consumer, user), nil
}

func NewConsumerUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &ConsumerUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}
