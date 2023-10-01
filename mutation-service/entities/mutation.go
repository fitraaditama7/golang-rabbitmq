package entities

import (
	"time"

	"github.com/gofrs/uuid"
)

type AccountMutation struct {
	ID              uuid.UUID  `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	AccountNumber   string     `gorm:"column:account_number"`
	TransactionTime time.Time  `gorm:"column:transaction_date"`
	TransactionCode string     `gorm:"column:transaction_code"`
	Amount          float64    `gorm:"column:amount"`
	CreatedAt       time.Time  `gorm:"column:created_at"`
	UpdatedAt       *time.Time `gorm:"column:updated_at;autoUpdateTime:milli;"`
}
