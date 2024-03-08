package main

import (
	fmt "fmt"

	proto "go_trial/gorest/microservice/encryptService/protofiles"

	micro "go-micro.dev/v4"
)

func main() {
	//Create a ner service. Optionally include some options here
	service := micro.NewService(micro.Name("encrypter"))

	//init will parse the command line flags

	service.Init()

	//Register handler
	proto.RegisterEncrypterHandler(service.Server(), new(Encrypter))

	//Run the server

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
