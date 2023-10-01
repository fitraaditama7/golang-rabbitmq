package repository

import (
	"account-service/entities"
	"context"
)

//go:generate mockgen -destination=../mocks/repository/customer_repository_mock.go -package=mock_repository . CustomerRepository
type AccountRepository interface {
	GetByNIK(ctx context.Context, nik string) (*entities.Account, error)
	GetByPhoneNumber(ctx context.Context, phoneNumber string) (*entities.Account, error)
	GetByAccountNumber(ctx context.Context, accountNumber string) (*entities.Account, error)
	Insert(ctx context.Context, account *entities.Account) error
	Update(ctx context.Context, account *entities.Account) error
}
