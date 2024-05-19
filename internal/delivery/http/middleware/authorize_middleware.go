package middleware

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/constant"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/consumer"
	"github.com/gofiber/fiber/v2"
)

func NewAuthorizeMiddleware(consumer *consumer.ConsumerUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		xCtx := ctx.UserContext()
		user := GetUser(ctx)

		if user.Role == constant.Admin {
			return ctx.Next()
		}

		_, err := consumer.Repository.Consumer.FindByUserID(consumer.DB.WithContext(xCtx), user.UserID)
		if err != nil {
			consumer.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		return ctx.Next()
	}
}
