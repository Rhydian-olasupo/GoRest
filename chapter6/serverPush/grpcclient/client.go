package main

import (
	pb "go_trial/gorest/chapter6/serverPush/protofiles"
	"io"
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address = "localhost:50051"
)

//Receieve Stream Lsitens to the Stream contents and use them

func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started Lsitening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _,%v", client, err)
	}

	//Listen to messages
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			//IF there are no monre messages to listen to, get ou tof the loop
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _,%v", client, err)
		}

		log.Printf("Status:%v, Operation: %v", response.Status, response.Description)
	}

}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMoneyTransactionClient(conn)

	// Prepare data. Get this from clients like Front-end or Android App
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// Contact the server and print out its response.
	ReceiveStream(client, &pb.TransactionRequest{From: from,
		To: to, Amount: amount})
}
