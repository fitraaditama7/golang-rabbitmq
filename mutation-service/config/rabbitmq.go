package config

import (
	"fmt"
	"mutation-service/utils/logger"

	"github.com/streadway/amqp"
)

func InitRabbitMQ() (*amqp.Connection, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		Config("RABBITMQ_USER"),
		Config("RABBITMQ_PASSWORD"),
		Config("RABBITMQ_HOST"),
		Config("RABBITMQ_PORT"),
	)
	conn, err := amqp.Dial(url)
	if err != nil {
		var l = logger.Log()
		l.Error().Err(err).Msg(err.Error())
		return nil, err
	}
	return conn, nil
}
