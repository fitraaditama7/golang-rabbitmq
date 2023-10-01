package service

import (
	"account-service/models"

	"github.com/gofiber/fiber/v2"
)

//go:generate mockgen -destination=../mocks/service/account_service_mock.go -package=mock_service . AccountService
type AccountService interface {
	Register(c *fiber.Ctx, request *models.AccountRegisterRequest) error
	CheckBalance(c *fiber.Ctx, accountNumber string) error
	Saving(c *fiber.Ctx, request *models.AccountSavingRequest) error
	Withdraw(c *fiber.Ctx, request *models.AccountWithdrawequest) error
}
