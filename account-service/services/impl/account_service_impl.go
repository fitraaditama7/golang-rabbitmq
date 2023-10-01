package service_impl

import (
	"account-service/constants"
	"account-service/exceptions"
	"account-service/mapper"
	"account-service/models"
	rabbitmq_outbound "account-service/outbound/rabbitmq"
	repository "account-service/repositories"
	service "account-service/services"
	"account-service/utils/logger"
	"account-service/utils/responses"
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v2"
)

type accountService struct {
	accountRepository repository.AccountRepository
	rabbitmqOutbound  rabbitmq_outbound.RabbitmqOutbound
}

func NewAccountService(accountRepository repository.AccountRepository, rabbitmqOutbound rabbitmq_outbound.RabbitmqOutbound) service.AccountService {
	return &accountService{
		accountRepository: accountRepository,
		rabbitmqOutbound:  rabbitmqOutbound,
	}
}

func (a accountService) Register(c *fiber.Ctx, request *models.AccountRegisterRequest) error {
	var l = logger.Log()
	var ctx = c.Context()

	accountNIKCheck, err := a.accountRepository.GetByNIK(ctx, request.NIK)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	if accountNIKCheck != nil {
		return exceptions.BadRequest(c, constants.ErrorNIKAlreadyExists)
	}

	accountPhoneNumberCheck, err := a.accountRepository.GetByPhoneNumber(ctx, request.PhoneNumber)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	if accountPhoneNumberCheck != nil {
		return exceptions.BadRequest(c, constants.ErrorPhoneNumberAlreadyExists)
	}

	account := mapper.AccountRegisterRequestToAccountEntity(request)

	err = a.accountRepository.Insert(ctx, &account)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	result := mapper.AccountEntityToAccountResponse(&account)
	return c.JSON(responses.Success(result))
}

func (a accountService) CheckBalance(c *fiber.Ctx, accountNumber string) error {
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

	result := mapper.AccountEntityToAccountResponse(account)
	return c.JSON(responses.Success(result))
}

func (a accountService) Saving(c *fiber.Ctx, request *models.AccountSavingRequest) error {
	var l = logger.Log()
	var ctx = c.Context()

	account, err := a.accountRepository.GetByAccountNumber(ctx, request.AccountNumber)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	if account == nil {
		return exceptions.DataNotFound(c, constants.ErrorDataNotFound)
	}

	now := time.Now()
	account.Balance = account.Balance + request.Amount
	account.UpdatedAt = &now

	err = a.accountRepository.Update(ctx, account)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	go func() {
		accountMutation := mapper.AccountMutationBuild(account.UpdatedAt, account.AccountNumber, "Saving", request.Amount)
		err = a.produceAccountMutationData(accountMutation)
		if err != nil {
			l.Error().Err(err).Msg(err.Error())
		}
	}()

	result := mapper.AccountEntityToAccountResponse(account)
	return c.JSON(responses.Success(result))
}

func (a accountService) Withdraw(c *fiber.Ctx, request *models.AccountWithdrawequest) error {
	var l = logger.Log()
	var ctx = c.Context()

	account, err := a.accountRepository.GetByAccountNumber(ctx, request.AccountNumber)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	if account == nil {
		return exceptions.DataNotFound(c, constants.ErrorDataNotFound)
	}

	if account.Balance < request.Amount {
		return exceptions.DataNotFound(c, constants.ErrorInvalidAmount)
	}

	now := time.Now()
	account.Balance = account.Balance - request.Amount
	account.UpdatedAt = &now

	err = a.accountRepository.Update(ctx, account)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.SystemError(c)
	}

	go func() {
		accountMutation := mapper.AccountMutationBuild(account.UpdatedAt, account.AccountNumber, "Withdraw", request.Amount)
		err = a.produceAccountMutationData(accountMutation)
		if err != nil {
			l.Error().Err(err).Msg(err.Error())
		}
	}()

	result := mapper.AccountEntityToAccountResponse(account)
	return c.JSON(responses.Success(result))
}

func (a accountService) produceAccountMutationData(data models.RabbitMQAccountMutationRequest) error {
	var l = logger.Log()

	b, err := json.Marshal(data)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return err
	}

	err = a.rabbitmqOutbound.Produce(b)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return err
	}
	return nil
}
