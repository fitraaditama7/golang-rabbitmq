package models

import "time"

type AccountMutationResponse struct {
	Name          string             `json:"nama"`
	AccountNumber string             `json:"no_rekening"`
	Mutation      []MutationResponse `json:"mutasi"`
}

type MutationResponse struct {
	TransactionTime *time.Time `json:"waktu"`
	TransactionCode string     `json:"kode_transaksi"`
	Amount          float64    `json:"nominal"`
}
