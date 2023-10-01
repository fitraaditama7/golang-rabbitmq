package models

type AccountWithdrawequest struct {
	AccountNumber string  `json:"nomor_rekening" validate:"required" example:"123459"`
	Amount        float64 `json:"nominal" validate:"required,numeric" example:"123"`
}
