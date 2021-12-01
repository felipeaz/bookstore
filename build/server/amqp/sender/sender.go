package sender

import (
	"bookstore/internal/app/logger"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	QueueConn    *amqp.Connection
	QueueChannel *amqp.Channel
	Queue        amqp.Queue
	QueueName    string
	Logger       logger.LogInterface
}

func CreateAMQP(host, queueName string, log logger.LogInterface) (*RabbitMQ, error) {
	connRabbitMQ, err := amqp.Dial(host)
	if err != nil {
		log.Error(err)
		return nil, nil
	}

	chRabbitMQ, err := connRabbitMQ.Channel()
	if err != nil {
		log.Error(err)
		return nil, nil
	}

	queue, err := chRabbitMQ.QueueDeclare(
		queueName, // queue name
		true,      // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		QueueConn:    connRabbitMQ,
		QueueChannel: chRabbitMQ,
		Queue:        queue,
	}, nil
}

func (q *RabbitMQ) PushMessage(msg []byte) error {
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	}
	err := q.QueueChannel.Publish(
		"",
		q.QueueName,
		false,
		false,
		message,
	)
	if err != nil {
		q.Logger.Error(err)
		return err
	}
	return nil
}
