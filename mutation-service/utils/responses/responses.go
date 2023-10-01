package responses

import "mutation-service/models"

func Error(errorMessage []string) models.ResponseError {
	return models.ResponseError{
		Error: []models.ResponseData{constructResponse(nil, errorMessage[0], errorMessage[1])},
	}
}

func Success(data any) models.ResponseData {
	return models.ResponseData{
		Code:    "00",
		Message: "SUCCESS",
		Data:    data,
	}
}

func constructResponse(data any, code, message string) models.ResponseData {
	return models.ResponseData{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
