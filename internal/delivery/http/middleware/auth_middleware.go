package middleware

import (
	"fmt"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

func NewAuth(userUseCase *user.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var token string
		xCtx := ctx.UserContext()
		request := &model.VerifyUserRequest{
			Token: ctx.Get("Authorization", "NOT_FOUND"),
		}

		parts := strings.Fields(request.Token)

		if len(parts) == 2 && parts[0] == "Bearer" {
			token = parts[1]
		} else {
			userUseCase.Log.Warnf("Invalid token format")
			return fiber.ErrUnauthorized
		}
		request.Token = token
		userUseCase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userUseCase.Verify(xCtx, request)
		if err != nil {
			userUseCase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		err = CheckTokenExpired(userUseCase, token)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		userUseCase.Log.Debugf("User : %+v", auth)

		authentic := model.Auth{
			Email:  auth.Email,
			UserID: auth.UserID,
			Role:   auth.Role,
		}

		ctx.Locals("auth", authentic)
		return ctx.Next()
	}
}

func CheckTokenExpired(useCase *user.UserUseCase, token string) error {
	jwtSecret := useCase.Config.JWT.Secret
	claims := &jwt.StandardClaims{}

	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtSecret), nil
	})

	if !parsedToken.Valid {
		useCase.Log.Warnf("Invalid token : %+v", err)
		return fiber.ErrUnauthorized
	}

	if claims.ExpiresAt < time.Now().Unix() {
		useCase.Log.Warnf("Token expired at: %v", claims.ExpiresAt)
		return fiber.ErrUnauthorized
	}

	if err != nil {
		useCase.Log.Warnf("Failed to parse JWT token : %+v", err)
		return fiber.ErrUnauthorized
	}

	return nil
}

func GetUser(ctx *fiber.Ctx) model.Auth {
	return ctx.Locals("auth").(model.Auth)
}
