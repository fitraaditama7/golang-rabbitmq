package models

type AccountRegisterRequest struct {
	Name        string `json:"nama" validate:"required" example:"John Doe"`
	NIK         string `json:"nik" validate:"required,number" example:"123xxxxx"`
	PhoneNumber string `json:"no_hp" validate:"required" example:"0812xxxxxxxx"`
}
