package mapper

import "account-service/models"

func RestAccountMutationDataResponseToAccountMutationResponse(mutationResponse models.RestAccountMutationDataResponse) models.MutationResponse {
	return models.MutationResponse{
		TransactionTime: mutationResponse.TransactionTime,
		TransactionCode: mutationResponse.TransactionCode,
		Amount:          mutationResponse.Amount,
	}
}
