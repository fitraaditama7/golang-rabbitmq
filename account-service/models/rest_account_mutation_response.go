package models

import "time"

type RestAccountMutationResponse struct {
	Message string                            `json:"message"`
	Data    []RestAccountMutationDataResponse `json:"data"`
}

type RestAccountMutationDataResponse struct {
	AccountNumber   string     `json:"no_rekening"`
	TransactionTime *time.Time `json:"waktu"`
	TransactionCode string     `json:"kode_transaksi"`
	Amount          float64    `json:"nominal"`
}
