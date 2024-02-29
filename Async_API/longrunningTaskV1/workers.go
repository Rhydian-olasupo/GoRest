package main

import (
	"encoding/json"
	"go_trial/gorest/Async_API/longrunningTaskV1/models"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type Workers struct {
	conn amqp.Connection
}

func (w *Workers) run() {
	log.Printf("Workers are booted up and running")
	channel, err := w.conn.Channel()
	handleError(err, "Fetching Channel Failed")

	defer channel.Close()

	jobQueue, err := channel.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	handleError(err, "Job queue fetch Failed")

	messages, err := channel.Consume(
		jobQueue.Name, //queue
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	go func() {
		for message := range messages {
			job := models.Job{}
			err = json.Unmarshal(message.Body, &job)
			log.Printf("Workers received a message from the quue: %s", job)
			handleError(err, "Unable to load queue messages")

			switch job.Type {
			case "A":
				w.dbWork(job)
			case "B":
				w.callbackWork(job)
			case "C":
				w.emailWork(job)
			}

		}
	}()
	defer w.conn.Close()
	wait := make(chan bool)
	<-wait //. Run long-running worker.
}

func (w *Workers) dbWork(job models.Job) {
	result := job.ExtraData.(map[string]interface{})
	log.Printf("Workers %s: extracting data....., JOB: %s", job.Type, result)
	time.Sleep(2 * time.Second)
	log.Printf("Workesrs %s: saving data to database...., JOB: %s", job.Type, job.ID)
}

func (w *Workers) callbackWork(job models.Job) {
	log.Printf("Worker %s: performing some long running process..., JOB: %s", job.Type, job.ID)
	time.Sleep(10 * time.Second)
	log.Printf("Worker %s: posting the data back to the given callback..., JOB: %s", job.Type, job.ID)
}

func (w *Workers) emailWork(job models.Job) {
	log.Printf("Worker %s: sending the email..., JOB: %s",
		job.Type, job.ID)
	time.Sleep(2 * time.Second)
	log.Printf("Worker %s: sent the email successfully, JOB: %s", job.Type, job.ID)
}
