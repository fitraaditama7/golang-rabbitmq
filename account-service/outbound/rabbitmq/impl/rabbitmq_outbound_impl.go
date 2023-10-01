package rabbitmq_outbound_impl

import (
	rabbitmq_outbound "account-service/outbound/rabbitmq"
	"account-service/utils/logger"

	"github.com/streadway/amqp"
)

type rabbitmqOutbound struct {
	conn *amqp.Connection
}

func NewRabbitmqOutbound(conn *amqp.Connection) rabbitmq_outbound.RabbitmqOutbound {
	return &rabbitmqOutbound{conn: conn}
}

func (r rabbitmqOutbound) Produce(body []byte) error {
	var l = logger.Log()
	ch, err := r.conn.Channel()
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"account-queue", //name
		false,           // durable
		false,           //delete when unused
		false,           // exclusive
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return err
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immadiate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	if err != nil {
		l.Error().Err(err).Msg(err.Error())
		return err
	}
	return nil
}
