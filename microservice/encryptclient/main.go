package main

import (
	"context"
	"fmt"

	proto "go_trial/gorest/microservice/encryptclient/proto/protofiles"

	micro "go-micro.dev/v4"
)

func main() {
	//Create client microservice
	service := micro.NewService(micro.Name("encrypter.client"))
	service.Init()

	//Create new encrypter service instance
	encrypter := proto.NewEncrypterService("encrypter", service.Client())

	//Call the encrypter service

	rsp, err := encrypter.Encrypt(context.TODO(), &proto.Request{
		Message: "I am a message",
		Key:     "111023043350789514532147",
	})

	if err != nil {
		fmt.Println(err)
	}

	//Print the message

	fmt.Println(rsp.Result)

	//Call decrypter
	rsp, err = encrypter.Decrypt(context.TODO(), &proto.Request{
		Message: rsp.Result,
		Key:     "111023043350789514532147",
	})

	if err != nil {
		fmt.Println(err)
	}

	//Print Response

	fmt.Println(rsp.Result)
}
