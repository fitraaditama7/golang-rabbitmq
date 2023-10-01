package main

import (
	"account-service/config"
	controllers_impl "account-service/controllers/impl"
	"account-service/database"
	_ "account-service/docs"
	rabbitmq_outbound_impl "account-service/outbound/rabbitmq/impl"
	rest_outbound_impl "account-service/outbound/rest/impl"
	repository_impl "account-service/repositories/impl"
	"account-service/routers"
	service_impl "account-service/services/impl"
	"account-service/utils/metric"
	"log"

	"github.com/gofiber/swagger"
	"github.com/prometheus/client_golang/prometheus"
)

// @title Fiber Swagger Example API
// @version 3.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:5300
// @BasePath /
// @schemes http

func main() {
	config.LoadConfig()

	db := database.Init()

	rabbitMQ, err := config.InitRabbitMQ()
	if err != nil {
		log.Fatal("cannot connect to rabbitmq")
	}
	defer rabbitMQ.Close()

	metrics := metric.SetupPrometheus()
	prometheus.MustRegister(metrics)

	accountRepository := repository_impl.NewAccountRepository(db)

	rabbitmqOutbound := rabbitmq_outbound_impl.NewRabbitmqOutbound(rabbitMQ)
	restAccountMutationOutbound := rest_outbound_impl.NewRestAccountMutationOutbound(config.Config("ACCOUNT_MUTATION_SERVICE_BASE_URL"))

	accountService := service_impl.NewAccountService(accountRepository, rabbitmqOutbound)
	restAccountMutationService := service_impl.NewAccountMutationService(accountRepository, restAccountMutationOutbound)

	accountController := controllers_impl.NewAccountController(accountService)
	accountMutationController := controllers_impl.NewAccountMutationController(restAccountMutationService)

	app := config.InitFiberApp()

	app.Get("/swagger/*", swagger.HandlerDefault)

	routers.SetupRoutes(app, accountController, accountMutationController)

	log.Fatal(app.Listen(":5100"))
}
