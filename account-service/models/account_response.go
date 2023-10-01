package models

import (
	"time"

	"github.com/google/uuid"
)

type AccountResponse struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"nama"`
	NIK           string     `json:"nik"`
	PhoneNumber   string     `json:"no_hp"`
	AccountNumber string     `json:"nomor_rekening"`
	Balance       float64    `json:"saldo"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
