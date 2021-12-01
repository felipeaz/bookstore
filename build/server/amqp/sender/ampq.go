package sender

import (
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	QueueChannel *amqp.Channel
	Queue        amqp.Queue
}

func CreateAMQP(host, queueName string) (*RabbitMQ, error) {
	connRabbitMQ, err := amqp.Dial(host)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	defer connRabbitMQ.Close()

	chRabbitMQ, err := connRabbitMQ.Channel()
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	defer chRabbitMQ.Close()

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
		QueueChannel: chRabbitMQ,
		Queue:        queue,
	}, nil
}
