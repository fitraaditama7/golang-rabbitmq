package mapper

import (
	"mutation-service/entities"
	"mutation-service/models"
)

func AccountMutationRequestToAccountMutationEntity(accountMutation models.AccountMutationRequest) entities.AccountMutation {
	return entities.AccountMutation{
		AccountNumber:   accountMutation.AccountNumber,
		TransactionTime: *accountMutation.TransactionTime,
		TransactionCode: accountMutation.TransactionCode,
		Amount:          accountMutation.Amount,
	}
}
