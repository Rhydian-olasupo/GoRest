package main

import "log"

const (
	queueName  string = "jobQueue"
	hostString string = "127.0.0.1:8000"
)

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
