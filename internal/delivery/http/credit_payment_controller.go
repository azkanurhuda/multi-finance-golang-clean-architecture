package http

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/credit_payment"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CreditPaymentController struct {
	Log     *logrus.Logger
	UseCase credit_payment.UseCase
}

func NewCreditPaymentController(logger *logrus.Logger, useCase credit_payment.UseCase) *CreditPaymentController {
	return &CreditPaymentController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *CreditPaymentController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateCreditPaymentRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create credit payment : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreditPaymentResponse]{Data: response})
}

func (c *CreditPaymentController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.GetByID(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to get credit payment : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreditPaymentResponse]{Data: response})
}

func (c *CreditPaymentController) GetList(ctx *fiber.Ctx) error {
	response, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to list credit payment : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.CreditPaymentResponse]{Data: response})
}

func (c *CreditPaymentController) DeleteByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.Delete(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to delete credit payment : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}
