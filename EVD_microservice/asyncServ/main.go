package main

import (
	"context"
	"log"
	"time"

	proto "go_trial/gorest/EVD_microservice/asyncServ/protofiles"

	micro "go-micro.dev/v4"
)

func main() {
	//Create a new service. Optionally add options here.

	service := micro.NewService(
		micro.Name("weather"),
	)

	p := micro.NewEvent("alerts", service.Client())

	go func() {
		for now := range time.Tick(15 * time.Second) {
			log.Println("Publishing weather alert to Topic: alerts")
			p.Publish(context.Background(), &proto.Event{
				City:        "Munich",
				Timestamp:   now.UTC().Unix(),
				Temperature: 2,
			})

		}

	}()

	//Run Server
	if err := service.Run(); err != nil {
		log.Println(err)
	}

}
