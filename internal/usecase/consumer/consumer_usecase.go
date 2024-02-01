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
	"strconv"
	"time"
)

type ConsumerUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
}

func NewConsumerUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository) UseCase {
	return &ConsumerUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
	}
}

func (c ConsumerUseCase) Create(ctx context.Context, request *model.CreateConsumerRequest, nik string) (*model.ConsumerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	data, err := c.Repository.Consumer.FindByNIK(tx, nik)
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
		NIK:          nik,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		PlaceOfBirth: request.PlaceOfBirth,
		DateOfBirth:  dateOfBirth,
		Salary:       request.Salary,
		IDCardPhoto:  idCardPhotoString,
		SelfiePhoto:  selfiePhotoString,
		CreatedAt:    time.Now().Local(),
		UpdatedAt:    time.Now().Local(),
	}

	if err := c.Repository.Consumer.Create(tx, consumer); err != nil {
		c.Log.Warnf("Failed create consumer to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	limit := 100000
	tempLimit := 0
	var rangeLimits = []int{100000, 300000, 200000}
	for i := 1; i <= 4; i++ {
		if i > 1 {
			tempLimit = rangeLimits[i-2]
		}
		limit = limit + tempLimit

		consumerLimit := &entity.ConsumerLimit{
			ID:        uuid.NewString(),
			NIK:       nik,
			Tenor:     strconv.Itoa(i),
			Limits:    int64(limit),
			CreatedAt: time.Now().Local(),
			UpdatedAt: time.Now().Local(),
		}

		if err := c.Repository.ConsumerLimit.Create(tx, consumerLimit); err != nil {
			c.Log.Warnf("Failed create consumer limit to database : %+v", err)
			return nil, fiber.ErrInternalServerError
		}
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.ConsumerToResponse(consumer), nil
}
