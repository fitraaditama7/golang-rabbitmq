package mapper

import (
	"mutation-service/entities"
	"mutation-service/models"
)

func AccountMutationEntityToAccountMutationResponse(accountMutation *entities.AccountMutation) models.AccountMutationResponse {
	return models.AccountMutationResponse{
		AccountNumber:   accountMutation.AccountNumber,
		TransactionTime: &accountMutation.TransactionTime,
		TransactionCode: accountMutation.TransactionCode,
		Amount:          accountMutation.Amount,
	}
}
