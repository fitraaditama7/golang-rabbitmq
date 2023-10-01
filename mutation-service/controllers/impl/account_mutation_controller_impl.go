package controllers_impl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mutation-service/constants"
	"mutation-service/controllers"
	"mutation-service/exceptions"
	"mutation-service/models"
	service "mutation-service/services"
	"mutation-service/utils/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/streadway/amqp"
	"github.com/thoas/go-funk"
)

type accountMutationController struct {
	accountMutationService service.AccountMutationService
}

func NewAccountMutationController(accountMutationService service.AccountMutationService) controllers.AccountMutationController {
	return &accountMutationController{
		accountMutationService: accountMutationService,
	}
}

func (a accountMutationController) GetByAccountNumber(c *fiber.Ctx) error {
	accountNumber := c.Params("no_rekening")

	if funk.IsEmpty(accountNumber) {
		return exceptions.BadRequest(c, constants.CustomErrorBadRequest(constants.ErrorEmptyAccountNumber[1]))
	}

	return a.accountMutationService.GetByAccountNumber(c, accountNumber)
}

func (a accountMutationController) ConsumeMessage(conn *amqp.Connection) {
	var l = logger.Log()
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"account-queue", // name
		false,           // durable
		false,           // delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal(err)
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			l.Info().Msg(fmt.Sprintf("Received a message: %s", d.Body))

			accountMutation := models.AccountMutationRequest{}
			err := json.Unmarshal(d.Body, &accountMutation)
			if err != nil {
				l.Error().Err(err).Msg(err.Error())
				continue
			}
			a.accountMutationService.Insert(context.Background(), accountMutation)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
