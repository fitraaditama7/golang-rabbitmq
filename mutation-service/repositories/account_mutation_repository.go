package repository

import (
	"context"
	"mutation-service/entities"
)

type AccountMutationRepository interface {
	GetByAccountNumber(ctx context.Context, accountNumber string) ([]*entities.AccountMutation, error)
	Create(ctx context.Context, accountMutation *entities.AccountMutation) error
}
