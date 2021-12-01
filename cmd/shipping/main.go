package main

import (
	"bookstore/build/server/amqp/consumer"
	_log "bookstore/infra/logger"
	"log"
	"os"
)

const (
	ServiceName = "Shipment Service"
)

func main() {
	logger := _log.NewLogger(os.Getenv("LOG_FILE"), ServiceName)
	queue, err := consumer.CreateAMQP(os.Getenv("AMQP_SERVER_URL"), os.Getenv("AMQP_QUEUE_NAME"), logger)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer queue.QueueConn.Close()
	defer queue.QueueChannel.Close()

	msgChan := make(chan bool)
	go func() {
		for message := range queue.Messages {
			log.Printf(" > Received Order for Shipping: %s\n", message.Body)
		}
	}()
	<-msgChan
}
