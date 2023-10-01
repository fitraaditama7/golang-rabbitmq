package controllers

import "github.com/gofiber/fiber/v2"

type AccountMutationController interface {
	GetAccountMutation(c *fiber.Ctx) error
}
