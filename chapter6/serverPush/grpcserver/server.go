package main

import (
	"fmt"
	pb "go_trial/gorest/chapter6/serverPush/protofiles"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMoneyTransactionServer
}

const (
	port      = ":50051"
	noOfSteps = 4
)

//MakeTransaction Impelements MoneyTransaction.MakeTransaction

func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Got request for money transfer...")
	log.Printf("Amount: %f From A/c:%s , To A/c:%s", in.Amount, in.From, in.To)
	//Send Stream Here
	for i := 0; i < noOfSteps; i++ {
		time.Sleep(time.Second * 2)
		//Once Tasks is complete, send the sucessful message back to the client
		if err := stream.Send(&pb.TransactionResponse{Status: "good", Step: int32(i), Description: fmt.Sprintf("Perfomring step %d", int32(i))}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}
	log.Printf("Successfully transferred amount $%v from %v to %v", in.Amount, in.From, in.To)
	log.Printf("Thank you for Using Our Service")
	return nil

}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("Could not connect to server")
	}
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
