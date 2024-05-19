package http

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http/middleware"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/credit"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type CreditController struct {
	Log     *logrus.Logger
	UseCase credit.UseCase
}

func NewCreditController(logger *logrus.Logger, useCase credit.UseCase) *CreditController {
	return &CreditController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *CreditController) Create(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	request := new(model.CreateCreditRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request, auth)
	if err != nil {
		c.Log.Warnf("Failed to create credit : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreditResponse]{Data: response})
}

func (c *CreditController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.GetByID(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to get credit : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CreditResponse]{Data: response})
}

func (c *CreditController) GetList(ctx *fiber.Ctx) error {
	response, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to list credit : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.CreditResponse]{Data: response})
}

func (c *CreditController) DeleteByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.Delete(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to delete credit : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}
