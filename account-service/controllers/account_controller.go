package controllers

import "github.com/gofiber/fiber/v2"

type AccountController interface {
	Register(c *fiber.Ctx) error
	CheckBalance(c *fiber.Ctx) error
	Saving(c *fiber.Ctx) error
	Withdraw(c *fiber.Ctx) error
}
