package controllers_impl

import (
	"account-service/constants"
	"account-service/controllers"
	"account-service/exceptions"
	"account-service/models"
	service "account-service/services"
	"account-service/utils/logger"
	"account-service/utils/validator"

	"github.com/gofiber/fiber/v2"
	"github.com/thoas/go-funk"
)

type accountController struct {
	accountService service.AccountService
}

func NewAccountController(accountService service.AccountService) controllers.AccountController {
	return &accountController{
		accountService: accountService,
	}
}

// Register Account
// @Summary     Register Account
// @Description Register Account
// @Tags        Account
// @Accept      json
// @Param       request-body   body   models.AccountRegisterRequest  true  "request-body"
// @Produce     json
// @Success     200 {object} models.ResponseData{data=models.AccountResponse}
// @failure     400 {object} models.ResponseError{error=models.ResponseData}
// @failure     500 {object} models.ResponseError{error=models.ResponseData}
// @Router      /daftar [post]
func (a accountController) Register(c *fiber.Ctx) error {
	l := logger.Log()

	request := new(models.AccountRegisterRequest)

	if err := c.BodyParser(request); err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.BadRequest(c, constants.CustomErrorBadRequest(err.Error()))
	}

	err := validator.Validate(request)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.BadRequest(c, constants.CustomErrorBadRequest(err.Error()))
	}

	return a.accountService.Register(c, request)
}

// Check Account Balance
// @Summary     Check account balance
// @Description Check account balance
// @Tags        Account
// @Accept      json
// @Param       no_rekening   path   string  true  "no_rekening"
// @Produce     json
// @Success     200 {object} models.ResponseData{data=models.AccountResponse}
// @failure     400 {object} models.ResponseError{error=models.ResponseData}
// @failure     500 {object} models.ResponseError{error=models.ResponseData}
// @Router      /saldo/{no_rekening} [get]
func (a accountController) CheckBalance(c *fiber.Ctx) error {
	accountNumber := c.Params("no_rekening")

	if funk.IsEmpty(accountNumber) {
		return exceptions.BadRequest(c, constants.CustomErrorBadRequest(constants.ErrorEmptyAccountNumber[1]))
	}

	return a.accountService.CheckBalance(c, accountNumber)
}

// Saving Account Balance
// @Summary     Saving to increase savings balance
// @Description Saving to increase savings balance
// @Tags        Account
// @Accept      json
// @Param       request-body   body   models.AccountSavingRequest  true  "request-body"
// @Produce     json
// @Success     200 {object} models.ResponseData{data=models.AccountResponse}
// @failure     400 {object} models.ResponseError{error=models.ResponseData}
// @failure     500 {object} models.ResponseError{error=models.ResponseData}
// @Router      /tabung [post]
func (a accountController) Saving(c *fiber.Ctx) error {
	l := logger.Log()

	request := new(models.AccountSavingRequest)

	if err := c.BodyParser(request); err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.BadRequest(c, constants.CustomErrorBadRequest(err.Error()))
	}

	err := validator.Validate(request)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.BadRequest(c, constants.CustomErrorBadRequest(err.Error()))
	}

	return a.accountService.Saving(c, request)
}

// Withdraw Account Balance
// @Summary     Withdraw to decrease savings balance
// @Description Withdraw to decrease savings balance
// @Tags        Account
// @Accept      json
// @Param       request-body   body   models.AccountWithdrawequest  true  "request-body"
// @Produce     json
// @Success     200 {object} models.ResponseData{data=models.AccountResponse}
// @failure     400 {object} models.ResponseError{error=models.ResponseData}
// @failure     500 {object} models.ResponseError{error=models.ResponseData}
// @Router      /tarik [post]
func (a accountController) Withdraw(c *fiber.Ctx) error {
	l := logger.Log()

	request := new(models.AccountWithdrawequest)

	if err := c.BodyParser(request); err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.BadRequest(c, constants.CustomErrorBadRequest(err.Error()))
	}

	err := validator.Validate(request)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return exceptions.BadRequest(c, constants.CustomErrorBadRequest(err.Error()))
	}

	return a.accountService.Withdraw(c, request)
}
