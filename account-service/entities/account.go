package entities

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID  `gorm:"column:id;primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name          string     `gorm:"column:name"`
	NIK           string     `gorm:"column:nik"`
	PhoneNumber   string     `gorm:"column:phone_number"`
	AccountNumber string     `gorm:"column:account_number;default:CONCAT('', LPAD(nextval('account_seq')::text, 15, '0'))"`
	Balance       float64    `gorm:"column:balance;default:0"`
	CreatedAt     time.Time  `gorm:"column:created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;autoUpdateTime:milli;"`
}
