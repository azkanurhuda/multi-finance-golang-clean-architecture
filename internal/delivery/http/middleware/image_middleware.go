package middleware

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/constant"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/user"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func NewImageMiddleware(user *user.UserUseCase) fiber.Handler {
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

		authentic := model.Auth{
			Email:  auth.Email,
			UserID: data.ID,
			Role:   data.Role,
		}

		if data.Role == constant.Admin {
			ctx.Locals("auth", authentic)
			return ctx.Next()
		} else {
			consumer, err := user.Repository.Consumer.FindByImage(user.DB.WithContext(xCtx), "id_card_photo", data.ID)
			if err != nil {
				user.Log.Warnf("Failed find image by id : %+v", err)
				return fiber.ErrUnauthorized
			}

			if consumer.UserID != data.ID {
				user.Log.Warnf("data id is not equal consumer user_id")
				return fiber.ErrForbidden
			}
		}

		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}
