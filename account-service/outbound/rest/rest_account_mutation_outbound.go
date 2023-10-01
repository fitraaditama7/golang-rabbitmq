package rest_outbound

import "account-service/models"

type RestAccountMutationOutbound interface {
	GetMutationByAccountNumber(accountNumber string) (*models.RestAccountMutationResponse, error)
}
