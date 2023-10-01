package controllers_impl

import (
	"github.com/gofiber/fiber/v2"
	"mutation-service/constants"
	"mutation-service/controllers"
	"mutation-service/exceptions"
	service "mutation-service/services"
	"mutation-service/utils/logger"
	"strconv"
)

type sampleController struct {
	sampleService service.SampleService
}

func NewSampleController(sampleService service.SampleService) controllers.SampleController {
	return &sampleController{
		sampleService: sampleService,
	}
}

func (s sampleController) FindAll(ctx *fiber.Ctx) error {
	return s.sampleService.FindAll(ctx)
}

func (s sampleController) FindByID(ctx *fiber.Ctx) error {
	l := logger.Log()
	l.Info().Msg("--- CancelPayment is called ---")

	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.BadRequest(ctx, constants.CustomErrorBadRequest(err.Error()))
	}
	return s.sampleService.FindByID(ctx, id)
}
