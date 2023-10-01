package models

import "time"

type AccountMutationRequest struct {
	AccountNumber   string     `json:"no_rekening"`
	TransactionTime *time.Time `json:"waktu"`
	TransactionCode string     `json:"kode_transaksi"`
	Amount          float64    `json:"nominal"`
}
