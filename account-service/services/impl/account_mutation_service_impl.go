package service_impl

import (
	"account-service/constants"
	"account-service/exceptions"
	"account-service/mapper"
	"account-service/models"
	rest_outbound "account-service/outbound/rest"
	repository "account-service/repositories"
	service "account-service/services"
	"account-service/utils/logger"
	"account-service/utils/responses"

	"github.com/gofiber/fiber/v2"
)

type mutationService struct {
	accountRepository repository.AccountRepository
	restOutbound      rest_outbound.RestAccountMutationOutbound
}

func NewAccountMutationService(accountRepository repository.AccountRepository, restOutbound rest_outbound.RestAccountMutationOutbound) service.AccountMutationService {
	return &mutationService{
		accountRepository: accountRepository,
		restOutbound:      restOutbound,
	}
}

func (a mutationService) CheckMutation(c *fiber.Ctx, accountNumber string) error {
	var l = logger.Log()
	var ctx = c.Context()

	account, err := a.accountRepository.GetByAccountNumber(ctx, accountNumber)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	if account == nil {
		return exceptions.DataNotFound(c, constants.ErrorDataNotFound)
	}

	restAccountMutations, err := a.restOutbound.GetMutationByAccountNumber(accountNumber)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	result := models.AccountMutationResponse{
		Name:          account.Name,
		AccountNumber: accountNumber,
	}
	for _, restAccountMutationData := range restAccountMutations.Data {
		mutationData := mapper.RestAccountMutationDataResponseToAccountMutationResponse(restAccountMutationData)
		result.Mutation = append(result.Mutation, mutationData)
	}

	return c.JSON(responses.Success(result))
}
