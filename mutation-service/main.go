package main

import (
	"log"
	"mutation-service/config"
	controllers_impl "mutation-service/controllers/impl"
	"mutation-service/database"
	repository_impl "mutation-service/repositories/impl"
	"mutation-service/routers"
	service_impl "mutation-service/services/impl"
	"mutation-service/utils/metric"

	"github.com/prometheus/client_golang/prometheus"
)

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

	accountMutationRepository := repository_impl.NewAccountMutationRepository(db)

	accountMutationService := service_impl.NewAccountMutationService(accountMutationRepository)

	accountMutationController := controllers_impl.NewAccountMutationController(accountMutationService)

	go func() {
		accountMutationController.ConsumeMessage(rabbitMQ)
	}()

	app := config.InitFiberApp()

	routers.SetupRoutes(app, accountMutationController)

	log.Fatal(app.Listen(":5200"))
}
