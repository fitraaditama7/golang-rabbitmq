package service

import (
	"context"
	"mutation-service/models"

	"github.com/gofiber/fiber/v2"
)

//go:generate mockgen -destination=../mocks/service/account_mutation_service_mock.go -package=mock_service . AccountMutationService
type AccountMutationService interface {
	GetByAccountNumber(c *fiber.Ctx, accountNumber string) error
	Insert(ctx context.Context, request models.AccountMutationRequest) error
}
