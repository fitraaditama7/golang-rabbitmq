package service

import "github.com/gofiber/fiber/v2"

type AccountMutationService interface {
	CheckMutation(c *fiber.Ctx, accountNumber string) error
}
