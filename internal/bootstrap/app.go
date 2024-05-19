package bootstrap

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/config"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http/middleware"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http/route"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/repository"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/asset"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/consumer"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/credit"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/credit_limit"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/credit_payment"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/merchant"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/transaction"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/user"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *config.Config
}

func Bootstrap(cfg *BootstrapConfig) {
	// setup repositories
	repo := repository.NewRepository(cfg.DB)

	// setup use cases
	userUseCase := user.NewUserUseCase(cfg.DB, cfg.Log, cfg.Validate, repo, cfg.Config)
	consumerUseCase := consumer.NewConsumerUseCase(cfg.DB, cfg.Log, cfg.Validate, repo)
	transactionUseCase := transaction.NewTransactionUseCase(cfg.DB, cfg.Log, cfg.Validate, repo)
	assetUseCase := asset.NewAssetUseCase(cfg.DB, cfg.Log, cfg.Validate, repo)
	creditUseCase := credit.NewCreditUseCase(cfg.DB, cfg.Log, cfg.Validate, repo)
	creditLimitUseCase := credit_limit.NewCreditLimitUseCase(cfg.DB, cfg.Log, cfg.Validate, repo)
	creditPaymentUseCase := credit_payment.NewCreditPaymentUseCase(cfg.DB, cfg.Log, cfg.Validate, repo)
	merchantUseCase := merchant.NewMerchantUseCase(cfg.DB, cfg.Log, cfg.Validate, repo)

	// setup controller
	userController := http.NewUserController(cfg.Log, userUseCase)
	consumerController := http.NewConsumerController(cfg.Log, consumerUseCase)
	transactionController := http.NewTransactionController(cfg.Log, transactionUseCase)
	assetController := http.NewAssetController(cfg.Log, assetUseCase)
	creditController := http.NewCreditController(cfg.Log, creditUseCase)
	creditLimitController := http.NewCreditLimitController(cfg.Log, creditLimitUseCase)
	creditPaymentController := http.NewCreditPaymentController(cfg.Log, creditPaymentUseCase)
	merchantController := http.NewMerchantController(cfg.Log, merchantUseCase)

	//setup middleware
	authMiddleware := middleware.NewAuth(&user.UserUseCase{
		DB:         cfg.DB,
		Log:        cfg.Log,
		Validate:   cfg.Validate,
		Repository: repo,
		Config:     cfg.Config,
	})
	adminMiddleware := middleware.NewAdminMiddleware(&user.UserUseCase{
		DB:         cfg.DB,
		Log:        cfg.Log,
		Validate:   cfg.Validate,
		Repository: repo,
		Config:     cfg.Config,
	})
	imageMiddleware := middleware.NewAdminMiddleware(&user.UserUseCase{
		DB:         cfg.DB,
		Log:        cfg.Log,
		Validate:   cfg.Validate,
		Repository: repo,
		Config:     cfg.Config,
	})

	authorizeMiddleware := middleware.NewAuthorizeMiddleware(&consumer.ConsumerUseCase{
		DB:         cfg.DB,
		Log:        cfg.Log,
		Validate:   cfg.Validate,
		Repository: repo,
	})

	routeConfig := route.Config{
		App:                     cfg.App,
		UserController:          userController,
		AuthMiddleware:          authMiddleware,
		AdminMiddleware:         adminMiddleware,
		ImageMiddleware:         imageMiddleware,
		AuthorizeMiddleware:     authorizeMiddleware,
		ConsumerController:      consumerController,
		TransactionController:   transactionController,
		AssetController:         assetController,
		CreditController:        creditController,
		CreditLimitController:   creditLimitController,
		CreditPaymentController: creditPaymentController,
		MerchantController:      merchantController,
	}

	routeConfig.Setup()
}
