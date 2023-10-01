package routers

import (
	"mutation-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, accountMutationController controllers.AccountMutationController) {
	app.Get("/mutasi/:no_rekening", accountMutationController.GetByAccountNumber)
}
