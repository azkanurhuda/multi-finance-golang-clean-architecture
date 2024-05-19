package http

import (
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/constant"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/delivery/http/middleware"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/model"
	"github.com/azkanurhuda/multi-finance-golang-clean-architecture/internal/usecase/consumer"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type ConsumerController struct {
	Log     *logrus.Logger
	UseCase consumer.UseCase
}

func NewConsumerController(logger *logrus.Logger, useCase consumer.UseCase) *ConsumerController {
	return &ConsumerController{
		Log:     logger,
		UseCase: useCase,
	}
}

func (c *ConsumerController) CreateConsumer(ctx *fiber.Ctx) error {
	auth := middleware.GetUser(ctx)
	request := new(model.CreateConsumerRequest)
	requestByAdmin := new(model.CreateConsumerByAdminRequest)
	var response *model.ConsumerResponse
	var err error

	if auth.Role == constant.Admin {
		err = ctx.BodyParser(requestByAdmin)
		if err != nil {
			c.Log.Warnf("Failed to parse request body : %+v", err)
			return fiber.ErrBadRequest
		}

		response, err = c.UseCase.CreateByAdmin(ctx.UserContext(), requestByAdmin)
		if err != nil {
			c.Log.Warnf("Failed to create consumer : %+v", err)
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
			c.Log.Warnf("Failed to create consumer : %+v", err)
			return err
		}
	}

	return ctx.JSON(model.WebResponse[*model.ConsumerResponse]{Data: response})
}
