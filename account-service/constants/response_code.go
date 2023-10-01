package constants

var ErrorServiceUnavailable = []string{"SERVICE_UNAVAILABLE", "Service unavailable"}
var ErrorInternalServer = []string{"INTERNAL_SERVER_ERROR", "Something wrong in our backend"}
var ErrorNIKAlreadyExists = []string{"BAD_REQUEST", "NIK already exists"}
var ErrorPhoneNumberAlreadyExists = []string{"BAD_REQUEST", "Phone number already exists"}
var ErrorInvalidAmount = []string{"BAD_REQUEST", "Current amount is less than withdraw amount"}
var ErrorEmptyAccountNumber = []string{"BAD_REQUEST", "Account number is empty"}
var ErrorBadRequest = []string{"BAD_REQUEST", "Bad Request"}
var ErrorDataNotFound = []string{"DATA_NOT_FOUND", "Data not found"}
var ErrorForbidden = []string{"FORBIDDEN", "Forbidden"}
var ErrorUnauthorized = []string{"UNAUTHORIZED", "Access Denied"}

func CustomErrorBadRequest(errorMessage string) []string {
	return []string{ErrorBadRequest[0], errorMessage}
}
