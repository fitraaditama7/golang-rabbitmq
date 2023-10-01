package responses

import "account-service/models"

func Error(errorMessage []string) models.ResponseError {
	return models.ResponseError{
		Error: []models.ResponseData{constructResponse(nil, "", errorMessage[0], errorMessage[1])},
	}
}

func Success(data any) models.ResponseData {
	return models.ResponseData{
		Message: "SUCCESS",
		Data:    data,
	}
}

func constructResponse(data any, code, message, remark string) models.ResponseData {
	return models.ResponseData{
		Remark:  remark,
		Message: message,
		Data:    data,
	}
}
