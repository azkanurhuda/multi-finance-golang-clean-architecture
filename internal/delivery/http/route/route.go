package route

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App                   *fiber.App
	UserController        *http.UserController
	AuthMiddleware        fiber.Handler
	ConsumerController    *http.ConsumerController
	TransactionController *http.TransactionController
}

func (c *RouteConfig) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupGuestRoute() {
	c.App.Post("/api/users", c.UserController.Register)
	c.App.Post("/api/users/_login", c.UserController.Login)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Delete("/api/users", c.UserController.Logout)
	c.App.Get("/api/users/_current", c.UserController.Current)

	c.App.Post("/api/consumers", c.ConsumerController.CreateConsumer)

	c.App.Post("/api/transactions", c.TransactionController.CreateTransaction)
}
