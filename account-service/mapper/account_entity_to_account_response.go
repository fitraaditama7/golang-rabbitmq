package mapper

import (
	"account-service/entities"
	"account-service/models"
)

func AccountEntityToAccountResponse(account *entities.Account) models.AccountResponse {
	return models.AccountResponse{
		ID:            account.ID,
		Name:          account.Name,
		NIK:           account.NIK,
		AccountNumber: account.AccountNumber,
		PhoneNumber:   account.PhoneNumber,
		Balance:       account.Balance,
		CreatedAt:     account.CreatedAt,
		UpdatedAt:     account.UpdatedAt,
	}
}
