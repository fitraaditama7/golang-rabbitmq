package service_impl

import (
	"context"
	"mutation-service/exceptions"
	"mutation-service/mapper"
	"mutation-service/models"
	repository "mutation-service/repositories"
	service "mutation-service/services"
	"mutation-service/utils/logger"
	"mutation-service/utils/responses"

	"github.com/gofiber/fiber/v2"
)

type accountMutation struct {
	accountMutationRepository repository.AccountMutationRepository
}

func NewAccountMutationService(accountMutationRepository repository.AccountMutationRepository) service.AccountMutationService {
	return &accountMutation{
		accountMutationRepository: accountMutationRepository,
	}
}

func (a accountMutation) GetByAccountNumber(c *fiber.Ctx, accountNumber string) error {
	var l = logger.Log()

	accountMutations, err := a.accountMutationRepository.GetByAccountNumber(c.Context(), accountNumber)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	result := []models.AccountMutationResponse{}
	for _, accountMutation := range accountMutations {
		result = append(result, mapper.AccountMutationEntityToAccountMutationResponse(accountMutation))
	}

	return c.JSON(responses.Success(result))
}

func (a accountMutation) Insert(ctx context.Context, request models.AccountMutationRequest) error {
	var l = logger.Log()

	accountMutation := mapper.AccountMutationRequestToAccountMutationEntity(request)
	err := a.accountMutationRepository.Create(ctx, &accountMutation)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return err
	}
	return nil
}
