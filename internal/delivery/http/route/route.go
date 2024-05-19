package route

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App                     *fiber.App
	UserController          *http.UserController
	AuthMiddleware          fiber.Handler
	AdminMiddleware         fiber.Handler
	ImageMiddleware         fiber.Handler
	AuthorizeMiddleware     fiber.Handler
	ConsumerController      *http.ConsumerController
	TransactionController   *http.TransactionController
	AssetController         *http.AssetController
	CreditController        *http.CreditController
	CreditLimitController   *http.CreditLimitController
	CreditPaymentController *http.CreditPaymentController
	MerchantController      *http.MerchantController
}

func (c *Config) Setup() {
	c.SetupGuestRoute()
	c.SetupAuthRoute()
	c.SetupAdminRoute()
}

func (c *Config) SetupGuestRoute() {
	api := c.App.Group("/api")

	// User
	api.Post("/user/register", c.UserController.Register)
	api.Post("/user/login", c.UserController.Login)

	// TODO Consumer

	// Merchant
	api.Get("/merchant/detail/:id", c.MerchantController.GetByID)
	api.Get("/merchant/list", c.MerchantController.GetList)

	// Asset
	api.Get("/asset/detail/:id", c.AssetController.GetByID)
	api.Get("/asset/list", c.AssetController.GetList)
}

func (c *Config) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	api := c.App.Group("/api")
	api.Delete("/user", c.UserController.Logout)
	api.Get("/user/current", c.UserController.Current)

	// Consumer
	api.Post("/consumer/create", c.ConsumerController.CreateConsumer)

	authorize := api.Use(c.AuthorizeMiddleware)
	// Credit
	authorize.Post("/credit/create", c.CreditController.Create)
	authorize.Get("/credit/detail/:id", c.CreditController.GetByID)
	authorize.Get("/credit/list", c.CreditController.GetList)

	// Credit Limit
	authorize.Get("/credit/limit/detail/:id", c.CreditLimitController.GetByID)
	authorize.Get("/credit/limit/list", c.CreditLimitController.GetList)

	// Credit Payment
	authorize.Post("/credit/payment/create", c.CreditPaymentController.Create)
	authorize.Get("/credit/payment/detail/:id", c.CreditPaymentController.GetByID)
	authorize.Get("/credit/payment/list", c.CreditPaymentController.GetList)

	// Transaction
	authorize.Post("/transaction/create", c.TransactionController.CreateTransaction)
}

func (c *Config) SetupAdminRoute() {
	api := c.App.Group("/api/admin", c.AdminMiddleware)

	// Merchant
	api.Post("/merchant/create", c.MerchantController.Create)
	api.Delete("/merchant/delete/:id", c.MerchantController.DeleteByID)

	// Asset
	api.Post("/asset/create", c.AssetController.Create)
	api.Delete("/asset/delete/:id", c.AssetController.DeleteByID)

	// Credit
	api.Delete("/credit/delete/:id", c.CreditController.DeleteByID)

	// Credit Limit
	api.Post("/credit/limit/create", c.CreditLimitController.Create)
	api.Delete("/credit/limit/delete/:id", c.CreditLimitController.DeleteByID)

	// Credit Payment
	api.Delete("/credit/payment/delete/:id", c.CreditPaymentController.DeleteByID)
}
