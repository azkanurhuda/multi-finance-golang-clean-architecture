package http

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/constant"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http/middleware"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/transaction"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type TransactionController struct {
	Log     *logrus.Logger
	UseCase transaction.UseCase
}

func NewTransactionController(logger *logrus.Logger, useCase transaction.UseCase) *TransactionController {
	return &TransactionController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *TransactionController) CreateTransaction(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	request := new(model.CreateTransactionRequest)
	requestAdmin := new(model.CreateTransactionByAdminRequest)
	var response *model.TransactionResponse
	var err error

	if auth.Role == constant.Admin {
		err = ctx.BodyParser(requestAdmin)
		if err != nil {
			c.Log.Warnf("Failed to parse request body : %+v", err)
			return fiber.ErrBadRequest
		}

		response, err = c.UseCase.CreateByAdmin(ctx.UserContext(), requestAdmin)
		if err != nil {
			c.Log.Warnf("Failed to create transaction : %+v", err)
			return err
		}
	} else {
		err = ctx.BodyParser(request)
		if err != nil {
			c.Log.Warnf("Failed to parse request body : %+v", err)
			return fiber.ErrBadRequest
		}

		response, err = c.UseCase.Create(ctx.UserContext(), request, auth)
		if err != nil {
			c.Log.Warnf("Failed to create transaction : %+v", err)
			return err
		}
	}

	return ctx.JSON(model.WebResponse[*model.TransactionResponse]{Data: response})
}
