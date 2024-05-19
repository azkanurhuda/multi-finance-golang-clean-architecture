package user

import (
	"context"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/config"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/constant"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/entity"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model/converter"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserUseCase struct {
	DB         *gorm.DB
	Log        *logrus.Logger
	Validate   *validator.Validate
	Repository *repository.Repository
	Config     *config.Config
}

func NewUserUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate, repo *repository.Repository, config *config.Config) UseCase {
	return &UserUseCase{
		DB:         db,
		Log:        logger,
		Validate:   validate,
		Repository: repo,
		Config:     config,
	}
}

func (c UserUseCase) Verify(ctx context.Context, request *model.VerifyUserRequest) (*model.Auth, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	data, err := c.Repository.User.FindByToken(tx, request.Token)
	if err != nil {
		c.Log.Warnf("Failed find user by token : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &model.Auth{
		Email:  data.Email,
		UserID: data.ID,
		Role:   data.Role,
	}, nil
}

func (c UserUseCase) Create(ctx context.Context, request *model.RegisterUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	err := c.Validate.Struct(request)
	if err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	data, _ := c.Repository.User.FindByEmail(tx, request.Email)
	if data != nil {
		c.Log.Warnf("User already exists : %+v", data)
		return nil, fiber.ErrConflict
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.Log.Warnf("Failed to generate bcrype hash : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	token, err := c.GenerataToken(request.Email)
	if err != nil {
		c.Log.Warnf("Failed generate token : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user := &entity.User{
		ID:             uuid.New().String(),
		Email:          request.Email,
		Role:           constant.User,
		Password:       string(password),
		Token:          token.Token,
		TokenExpiredAt: token.TokenExpiredAt,
		CreatedAt:      time.Now().Local(),
		UpdatedAt:      time.Now().Local(),
	}

	if err := c.Repository.User.Create(tx, user); err != nil {
		c.Log.Warnf("Failed create user to database : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToTokenResponse(user), nil
}

func (c UserUseCase) Login(ctx context.Context, request *model.LoginUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body  : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	user := new(entity.User)
	data, err := c.Repository.User.FindByEmail(tx, request.Email)
	if data == nil {
		c.Log.Warnf("User not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(request.Password)); err != nil {
		c.Log.Warnf("Failed to compare user password with bcrype hash : %+v", err)
		return nil, fiber.ErrUnauthorized
	}

	token, err := c.GenerataToken(data.Email)
	if err != nil {
		c.Log.Warnf("Failed generate token : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	user.Token = token.Token
	user.TokenExpiredAt = token.TokenExpiredAt
	if err := c.Repository.User.UpdateTokenByEmail(tx, user, request.Email); err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToTokenResponse(user), nil
}

func (c UserUseCase) GenerataToken(email string) (*entity.User, error) {
	// Generate JWT token
	expirationTime := time.Now().Add(3 * time.Hour)
	claims := &jwt.StandardClaims{
		Subject:   email,
		ExpiresAt: expirationTime.Unix(),
	}

	jwtSecret := c.Config.JWT.Secret

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.Log.Warnf("Failed to sign JWT token : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return &entity.User{
		Token:          tokenString,
		TokenExpiredAt: &expirationTime,
	}, nil
}

func (c UserUseCase) Current(ctx context.Context, request *model.GetUserRequest) (*model.UserResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return nil, fiber.ErrBadRequest
	}

	data, err := c.Repository.User.FindByEmail(tx, request.Email)
	if data == nil {
		c.Log.Warnf("User not exists : %+v", err)
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return nil, fiber.ErrInternalServerError
	}

	return converter.UserToResponse(data), nil
}

func (c UserUseCase) Logout(ctx context.Context, request *model.LogoutUserRequest) (bool, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.Warnf("Invalid request body : %+v", err)
		return false, fiber.ErrBadRequest
	}

	user := new(entity.User)
	data, err := c.Repository.User.FindByEmail(tx, request.Email)
	if data == nil {
		c.Log.Warnf("User not exists : %+v", err)
		return false, fiber.ErrNotFound
	}

	user.Token = ""
	user.TokenExpiredAt = nil

	if err := c.Repository.User.UpdateTokenByEmail(tx, user, data.Email); err != nil {
		c.Log.Warnf("Failed save user : %+v", err)
		return false, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Warnf("Failed commit transaction : %+v", err)
		return false, fiber.ErrInternalServerError
	}

	return true, nil
}
