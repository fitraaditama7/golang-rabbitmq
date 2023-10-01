package service_impl

import (
	"github.com/gofiber/fiber/v2"
	"mutation-service/constants"
	"mutation-service/exceptions"
	"mutation-service/mapper"
	"mutation-service/models"
	repository "mutation-service/repositories"
	service "mutation-service/services"
	"mutation-service/utils/logger"
	"mutation-service/utils/responses"
)

type sampleService struct {
	sampleRepository repository.SampleRepository
}

func NewSampleService(sampleRepository repository.SampleRepository) service.SampleService {
	return &sampleService{
		sampleRepository: sampleRepository,
	}
}

func (s sampleService) FindAll(ctx *fiber.Ctx) error {
	var l = logger.Log()
	var err error

	samples, err := s.sampleRepository.FindAll(ctx.Context())
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(ctx)
	}

	modelSamples := []models.SampleResponse{}
	for _, sample := range samples {
		modelSample := mapper.SampleEntityToModel(sample)
		modelSamples = append(modelSamples, modelSample)
	}

	result := responses.Success(modelSamples)
	return ctx.JSON(result)
}

func (s sampleService) FindByID(ctx *fiber.Ctx, id int64) error {
	var l = logger.Log()
	var err error

	sample, err := s.sampleRepository.FindByID(ctx.Context(), id)
	if sample == nil {
		return exceptions.DataNotFound(ctx, constants.ErrorDataNotFound)
	}
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(ctx)
	}

	modelSample := mapper.SampleEntityToModel(sample)

	result := responses.Success(modelSample)
	return ctx.JSON(result)
}
