package constants

var ErrorServiceUnavailable = []string{"SERVICE_UNAVAILABLE", "Service unavailable"}
var ErrorInternalServer = []string{"INTERNAL_SERVER_ERROR", "Something wrong in our backend"}
var ErrorBadRequest = []string{"BAD_REQUEST", "Bad Request"}
var ErrorDataNotFound = []string{"DATA_NOT_FOUND", "Data not found"}
var ErrorForbidden = []string{"FORBIDDEN", "Forbidden"}
var ErrorUnauthorized = []string{"UNAUTHORIZED", "Access Denied"}
var ErrorEmptyAccountNumber = []string{"BAD_REQUEST", "Account number is empty"}

func CustomErrorBadRequest(errorMessage string) []string {
	return []string{ErrorBadRequest[0], errorMessage}
}
