package rest_outbound_impl

import (
	"account-service/models"
	rest_outbound "account-service/outbound/rest"
	"account-service/utils/logger"
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type restAccountMutationOutbound struct {
	baseURL string
	client  *fiber.Client
}

func NewRestAccountMutationOutbound(baseURL string) rest_outbound.RestAccountMutationOutbound {
	client := fiber.AcquireClient()

	return &restAccountMutationOutbound{
		baseURL: baseURL,
		client:  client,
	}
}

func (r restAccountMutationOutbound) GetMutationByAccountNumber(accountNumber string) (*models.RestAccountMutationResponse, error) {
	var l = logger.Log()

	var path = fmt.Sprintf("%s/mutasi/%s", r.baseURL, accountNumber)
	fmt.Println(path)
	var agent = r.client.Get(path)
	var responseBody models.RestAccountMutationResponse
	code, body, errs := agent.Struct(&responseBody)
	if len(errs) != 0 {
		l.Error().Err(errs[0]).Msg(errs[0].Error())
		return nil, errs[0]
	}

	if code != fiber.StatusOK {
		err := errors.New("outbound error")
		l.Error().Err(err).Msg(string(body))
		return nil, err
	}

	return &responseBody, nil
}
