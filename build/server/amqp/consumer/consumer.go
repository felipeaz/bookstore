package consumer

import (
	"bookstore/internal/app/logger"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	QueueConn    *amqp.Connection
	QueueChannel *amqp.Channel
	Messages     <-chan amqp.Delivery
	QueueName    string
	Logger       logger.LogInterface
}

func CreateAMQP(host, queueName string, log logger.LogInterface) (*RabbitMQ, error) {
	connRabbitMQ, err := amqp.Dial(host)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	chRabbitMQ, err := connRabbitMQ.Channel()
	if err != nil {
		log.Error(err)
		return nil, err
	}

	queue, err := chRabbitMQ.Consume(
		queueName, // queue name
		"",
		true,  // durable
		false, // auto delete
		false, // exclusive
		false, // no wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		QueueConn:    connRabbitMQ,
		QueueChannel: chRabbitMQ,
		Messages:     queue,
	}, nil
}
