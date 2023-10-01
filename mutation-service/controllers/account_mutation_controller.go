package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
)

type AccountMutationController interface {
	GetByAccountNumber(c *fiber.Ctx) error
	ConsumeMessage(conn *amqp.Connection)
}
