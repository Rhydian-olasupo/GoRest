package main

import (
	"log"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s ", err, msg)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Dialing failed to rabbitMQ ")

	defer conn.Close()

	channel, err := conn.Channel()
	handleError(err, "Fetching connection failed")

	defer channel.Close()

	testQueue, err := channel.QueueDeclare(
		"test",
		false,
		false,
		false,
		false,
		nil,
	)

	handleError(err, "Queue creation Failed")

	messages, err := channel.Consume(
		testQueue.Name,
		"", //consumer
		true,
		false,
		false,
		false,
		nil,
	)
	handleError(err, "Failed to register a consumer")

	go func() {
		for messages := range messages {
			log.Printf("Received a message from the queue %s", messages.Body)
		}
	}()

	log.Println("Worker has started")
	wait := make(chan bool)

	<-wait

}
