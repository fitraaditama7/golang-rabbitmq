package mapper

import (
	"account-service/models"
	"strings"
	"time"
)

func AccountMutationBuild(transactionTime *time.Time, accountNumber string, transactionType string, amount float64) models.RabbitMQAccountMutationRequest {
	var transactionCode = "C"
	if strings.ToLower(transactionType) == "withdraw" {
		transactionType = "D"
	}

	return models.RabbitMQAccountMutationRequest{
		AccountNumber:   accountNumber,
		TransactionTime: transactionTime,
		TransactionCode: transactionCode,
		Amount:          amount,
	}
}
