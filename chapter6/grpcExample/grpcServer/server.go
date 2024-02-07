package main

import (
	pb "go_trial/gorest/chapter6/grpcExample/protofiles"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//Server is used to create MoneyTransactionServer

type server struct {
	pb.UnimplementedMoneyTransactionServer
}

// MakeTransaction implements MoneyTransactionServer.MakeTransaction
func (s *server) MakeTransaction(ctx context.Context, in *pb.TranscationRequest) (*pb.TranscationResponse, error) {
	//Use in.Amount , in.From, in.To to perform transaction logic
	return &pb.TranscationResponse{Confirmation: true}, nil
}

const (
	port = ":50051"
)

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
