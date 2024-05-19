package bootstrap

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/config"
	"github.com/gofiber/fiber/v2"
)

func NewFiber(config *config.Config) *fiber.App {
	var app = fiber.New(fiber.Config{
		AppName:      config.App.Name,
		ErrorHandler: NewErrorHandler(),
		Prefork:      config.Server.Prefork,
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
