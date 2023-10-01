package routers

import (
	"account-service/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, accountController controllers.AccountController, accountMutationController controllers.AccountMutationController) {
	app.Post("/daftar", accountController.Register)
	app.Get("/saldo/:no_rekening", accountController.CheckBalance)
	app.Post("/tabung", accountController.Saving)
	app.Post("/tarik", accountController.Withdraw)

	app.Get("/mutasi/:no_rekening", accountMutationController.GetAccountMutation)
}
