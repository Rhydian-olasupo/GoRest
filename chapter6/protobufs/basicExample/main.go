package main

import (
	"fmt"

	pb "go_trial/gorest/chapter6/protobufs/protofiles/protofiles"

	"google.golang.org/protobuf/proto"
)

func main() {
	p := &pb.Person{
		Id:   1234,
		Name: "Roger D",
		Mail: "rf@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file:", p)
	fmt.Println("Marshalled proto data: ", body)
	fmt.Println("Unmarshalled struct: ", p1)

}
