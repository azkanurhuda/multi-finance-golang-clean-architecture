package middleware

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/constant"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func NewAdminMiddleware(user *user.UserUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		xCtx := ctx.UserContext()
		var token string
		request := &model.VerifyUserRequest{
			Token: ctx.Get("Authorization", "NOT_FOUND"),
		}

		parts := strings.Fields(request.Token)

		if len(parts) == 2 && parts[0] == "Bearer" {
			token = parts[1]
		} else {
			user.Log.Warnf("Invalid token format")
			return fiber.ErrUnauthorized
		}
		request.Token = token

		auth, err := user.Verify(xCtx, request)
		if err != nil {
			user.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		err = CheckTokenExpired(user, token)
		if err != nil {
			return fiber.ErrUnauthorized
		}

		data, err := user.Repository.User.FindByEmail(user.DB.WithContext(xCtx), auth.Email)
		if err != nil {
			user.Log.Warnf("Failed find user by email : %+v", err)
			return fiber.ErrUnauthorized
		}

		if data.Role != constant.Admin {
			return fiber.ErrForbidden
		}

		authentic := model.Auth{
			Email:  auth.Email,
			UserID: data.ID,
			Role:   data.Role,
		}

		ctx.Locals("auth", authentic)
		return ctx.Next()
	}
}
