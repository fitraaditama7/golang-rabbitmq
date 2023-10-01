package mapper

import (
	"account-service/entities"
	"account-service/models"
)

func AccountRegisterRequestToAccountEntity(request *models.AccountRegisterRequest) entities.Account {
	return entities.Account{
		Name:        request.Name,
		NIK:         request.NIK,
		PhoneNumber: request.PhoneNumber,
	}
}
