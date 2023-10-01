package repository_impl

import (
	"context"
	"mutation-service/entities"
	repository "mutation-service/repositories"

	"gorm.io/gorm"
)

type accountMutationRepository struct {
	DB *gorm.DB
}

func NewAccountMutationRepository(DB *gorm.DB) repository.AccountMutationRepository {
	return &accountMutationRepository{DB: DB}
}

func (a accountMutationRepository) GetByAccountNumber(ctx context.Context, accountNumber string) ([]*entities.AccountMutation, error) {
	var accountMutations []*entities.AccountMutation

	result := a.DB.WithContext(ctx).Order("transaction_date desc").Find(&accountMutations)
	if result.Error != nil {
		return nil, result.Error
	}

	return accountMutations, nil
}

func (a accountMutationRepository) Create(ctx context.Context, accountMutation *entities.AccountMutation) error {
	return a.DB.WithContext(ctx).Create(accountMutation).Error
}
