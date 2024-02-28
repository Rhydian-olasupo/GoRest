package main

import (
	_ "fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s :%s ", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	handleError(err, "Dialing failed to RabbitMq broker")
	defer conn.Close()

	channel, err := conn.Channel()
	handleError(err, "Fetching Channel Failed")

	defer channel.Close()

	testQueue, err := channel.QueueDeclare(
		"test", //Name of the Queue
		false,  // Persistency
		false,  //Delete message when unused
		false,  //Exclusivity
		false,  //No waiting time
		nil,    // Extra args
	)

	handleError(err, "Queue creation Failed")

	serverTime := time.Now()
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(serverTime.String()),
	}

	err = channel.Publish(
		"",
		testQueue.Name, //rooting key(@Queue)
		false,
		false,
		message,
	)

	handleError(err, "Failed to publish message")
	log.Println("Successfully published a message to the queue")

}
