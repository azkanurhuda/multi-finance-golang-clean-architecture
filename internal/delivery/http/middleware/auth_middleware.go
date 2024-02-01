package middleware

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func NewAuth(userUseCase *user.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var token string
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

		auth, err := userUseCase.Verify(ctx.UserContext(), request)
		if err != nil {
			userUseCase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}
		userUseCase.Log.Debugf("User : %+v", auth.NIK)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *model.Auth {
	return ctx.Locals("auth").(*model.Auth)
}
