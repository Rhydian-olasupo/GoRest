package main

import (
	"context"
	"log"

	proto "go_trial/gorest/EVD_microservice/asyncClient/protofiles"

	micro "go-micro.dev/v4"
)

// ProcessEvent processes a weather alert
func ProcessEvent(ctx context.Context, event *proto.Event) error {
	log.Println("Got alert:", event)
	return nil
}

func main() {
	//Create a new service
	service := micro.NewService(micro.Name("weather_client"))

	//Initialize the client and parse the command line flags

	service.Init()

	micro.RegisterSubscriber("alerts", service.Server(), ProcessEvent)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}