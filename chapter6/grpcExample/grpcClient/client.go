package main

import (
	pb "go_trial/gorest/chapter6/grpcExample/protofiles"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

func main() {
	// Setup a connection to the server

	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to dial %v", err)
	}

	//Create a client
	c := pb.NewMoneyTransactionClient(conn)

	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// Make server request.
	r, err := c.MakeTransaction(context.Background(), &pb.TranscationRequest{From: from, To: to, Amount: amount})

	if err != nil {
		log.Fatalf("Failed to make transaction: %v", err)
	}

	log.Printf("Transaction confirmed:%v", r)

}
