package service

import "github.com/gofiber/fiber/v2"

//go:generate mockgen -destination=../mocks/service/sample_service_mock.go -package=mock_service . SampleService
type SampleService interface {
	FindAll(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx, id int64) error
}
