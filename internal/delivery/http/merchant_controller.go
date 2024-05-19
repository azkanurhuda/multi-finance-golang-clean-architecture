package http

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/merchant"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type MerchantController struct {
	Log     *logrus.Logger
	UseCase merchant.UseCase
}

func NewMerchantController(logger *logrus.Logger, useCase merchant.UseCase) *MerchantController {
	return &MerchantController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *MerchantController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateMerchantRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create merchant : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.MerchantResponse]{Data: response})
}

func (c *MerchantController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.GetByID(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to get merchant : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.MerchantResponse]{Data: response})
}

func (c *MerchantController) GetList(ctx *fiber.Ctx) error {
	response, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to list merchant : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.MerchantResponse]{Data: response})
}

func (c *MerchantController) DeleteByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.Delete(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to delete merchant : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}
