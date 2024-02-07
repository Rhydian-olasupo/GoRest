package main

import (
	"fmt"

	pb "go_trial/gorest/chapter6/protobufs/protofiles/protofiles"

	"encoding/json"
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

	body, _ := json.Marshal(p)
	fmt.Println(string(body))

}
