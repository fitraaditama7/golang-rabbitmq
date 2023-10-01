package models

type ResponseData struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ResponseError struct {
	Error []ResponseData `json:"error"`
}
