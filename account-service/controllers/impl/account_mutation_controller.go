package controllers_impl

import (
	"account-service/constants"
	"account-service/controllers"
	"account-service/exceptions"
	service "account-service/services"

	"github.com/gofiber/fiber/v2"
	"github.com/thoas/go-funk"
)

type accountMutationController struct {
	accountMutationService service.AccountMutationService
}

func NewAccountMutationController(accountMutationService service.AccountMutationService) controllers.AccountMutationController {
	return &accountMutationController{
		accountMutationService: accountMutationService,
	}
}

// Get Account Mutation
// @Summary     Get Account Mutation
// @Description Get Account Mutation
// @Tags        Mutation
// @Accept      json
// @Param       no_rekening   path   string  true  "no_rekening"
// @Produce     json
// @Success     200 {object} models.ResponseData{data=models.AccountMutationResponse}
// @failure     400 {object} models.ResponseError{error=models.ResponseData}
// @failure     500 {object} models.ResponseError{error=models.ResponseData}
// @Router      /mutasi/{no_rekening} [get]
func (a accountMutationController) GetAccountMutation(c *fiber.Ctx) error {
	accountNumber := c.Params("no_rekening")

	if funk.IsEmpty(accountNumber) {
		return exceptions.BadRequest(c, constants.CustomErrorBadRequest(constants.ErrorEmptyAccountNumber[1]))
	}

	return a.accountMutationService.CheckMutation(c, accountNumber)
}
