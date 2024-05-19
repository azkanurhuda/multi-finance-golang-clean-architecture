package http

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/asset"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type AssetController struct {
	Log     *logrus.Logger
	UseCase asset.UseCase
}

func NewAssetController(logger *logrus.Logger, useCase asset.UseCase) *AssetController {
	return &AssetController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *AssetController) Create(ctx *fiber.Ctx) error {
	request := new(model.CreateAssetRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warnf("Failed to parse request body : %+v", err)
		return fiber.ErrBadRequest
	}

	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.Warnf("Failed to create asset : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AssetResponse]{Data: response})
}

func (c *AssetController) GetByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.GetByID(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to get asset : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[*model.AssetResponse]{Data: response})
}

func (c *AssetController) GetList(ctx *fiber.Ctx) error {
	response, err := c.UseCase.List(ctx.UserContext())
	if err != nil {
		c.Log.Warnf("Failed to list asset : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[[]model.AssetResponse]{Data: response})
}

func (c *AssetController) DeleteByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	response, err := c.UseCase.Delete(ctx.UserContext(), id)
	if err != nil {
		c.Log.Warnf("Failed to delete asset : %+v", err)
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: response})
}
