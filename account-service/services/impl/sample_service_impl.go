package service_impl

import (
	"account-service/constants"
	"account-service/exceptions"
	"account-service/mapper"
	"account-service/models"
	repository "account-service/repositories"
	service "account-service/services"
	"account-service/utils/logger"
	"account-service/utils/responses"
	"github.com/gofiber/fiber/v2"
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
