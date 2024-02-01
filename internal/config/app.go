package config

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http/middleware"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http/route"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/consumer"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/transaction"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	repo := repository.NewRepository(config.DB)

	// setup use cases
	userUseCase := user.NewUserUseCase(config.DB, config.Log, config.Validate, repo)
	consumerUseCase := consumer.NewConsumerUseCase(config.DB, config.Log, config.Validate, repo)
	transactionUseCase := transaction.NewTransactionUseCase(config.DB, config.Log, config.Validate, repo)

	// setup controller
	userController := http.NewUserController(config.Log, userUseCase)
	consumerController := http.NewConsumerController(config.Log, consumerUseCase)
	transactionController := http.NewTransactionController(config.Log, transactionUseCase)

	//setup middleware
	authMiddleware := middleware.NewAuth(&user.UserUseCase{
		DB:         config.DB,
		Log:        config.Log,
		Validate:   config.Validate,
		Repository: repo,
	})

	routeConfig := route.RouteConfig{
		App:                   config.App,
		UserController:        userController,
		AuthMiddleware:        authMiddleware,
		ConsumerController:    consumerController,
		TransactionController: transactionController,
	}

	routeConfig.Setup()
}
