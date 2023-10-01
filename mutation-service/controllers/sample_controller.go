package controllers

import (
	"github.com/gofiber/fiber/v2"
)

type SampleController interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
}
