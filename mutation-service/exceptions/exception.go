package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"mutation-service/constants"
	"mutation-service/utils/responses"
)

func BadRequest(ctx *fiber.Ctx, body []string) error {
	resp := responses.Error(body)
	return ctx.Status(fiber.StatusBadRequest).JSON(resp)
}

func Forbidden(ctx *fiber.Ctx, body []string) error {
	resp := responses.Error(body)
	return ctx.Status(fiber.StatusForbidden).JSON(resp)
}

func SystemError(ctx *fiber.Ctx) error {
	resp := responses.Error(constants.ErrorInternalServer)
	return ctx.Status(fiber.StatusInternalServerError).JSON(resp)
}

func DataNotFound(ctx *fiber.Ctx, body []string) error {
	resp := responses.Error(body)
	return ctx.Status(fiber.StatusForbidden).JSON(resp)
}
