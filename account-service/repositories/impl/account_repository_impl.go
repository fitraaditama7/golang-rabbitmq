package repository_impl

import (
	"account-service/entities"
	repository "account-service/repositories"
	"context"

	"gorm.io/gorm"
)

type accountRepository struct {
	DB *gorm.DB
}

func NewAccountRepository(DB *gorm.DB) repository.AccountRepository {
	return &accountRepository{DB: DB}
}

func (c *accountRepository) GetByNIK(ctx context.Context, nik string) (*entities.Account, error) {
	var account entities.Account
	result := c.DB.WithContext(ctx).Where("nik=?", nik).First(&account)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &account, nil
}

func (c *accountRepository) GetByPhoneNumber(ctx context.Context, phoneNumber string) (*entities.Account, error) {
	var account entities.Account
	result := c.DB.WithContext(ctx).Where("phone_number=?", phoneNumber).First(&account)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &account, nil
}

func (c *accountRepository) GetByAccountNumber(ctx context.Context, accountNumber string) (*entities.Account, error) {
	var account entities.Account
	result := c.DB.WithContext(ctx).Where("account_number=?", accountNumber).First(&account)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &account, nil
}

func (c *accountRepository) Insert(ctx context.Context, account *entities.Account) error {
	return c.DB.WithContext(ctx).Create(account).Error
}

func (c *accountRepository) Update(ctx context.Context, account *entities.Account) error {
	return c.DB.WithContext(ctx).Save(&account).Error
}
