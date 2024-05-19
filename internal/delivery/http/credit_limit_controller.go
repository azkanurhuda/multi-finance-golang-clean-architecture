package http

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/credit_limit"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CreditLimitController struct {
	Log     *logrus.Logger
	UseCase credit_limit.UseCase
}

func NewCreditLimitController(logger *logrus.Logger, useCase credit_limit.UseCase) *CreditLimitController {
	return &CreditLimitController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *CreditLimitController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateCreditLimitRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create credit limit : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreditLimitResponse]{Data: response})
}

func (c *CreditLimitController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.GetByID(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to get credit limit : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreditLimitResponse]{Data: response})
}

func (c *CreditLimitController) GetList(ctx *fiber.Ctx) error {
	response, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to list credit limit : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.CreditLimitResponse]{Data: response})
}

func (c *CreditLimitController) DeleteByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.Delete(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to delete credit limit : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}
