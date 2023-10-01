package models

type ResponseData struct {
	Message string `json:"message,omitempty" example:"SUCCESS/ERROR"`
	Remark  string `json:"remark,omitempty" example:"SUCCESS/ERROR"`
	Data    any    `json:"data,omitempty"`
}

type ResponseError struct {
	Error []ResponseData `json:"error"`
}
