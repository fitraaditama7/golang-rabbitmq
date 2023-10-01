package rabbitmq_outbound

type RabbitmqOutbound interface {
	Produce(body []byte) error
}
