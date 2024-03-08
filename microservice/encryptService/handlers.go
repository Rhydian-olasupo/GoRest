package main

import (
	"context"

	proto "go_trial/gorest/microservice/encryptService/protofiles"
)

//Encrypter holds the information about methods

type Encrypter struct{}

//Encrypt converts a message into cipher and returns response

func (g *Encrypter) Encrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Result = EncryptString(req.Key, req.Message)
	return nil
}

// Decrypt converts the cipher into original text and returns the response

func (g *Encrypter) Decrypt(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	rsp.Result = DecryptString(req.Key, req.Message)
	return nil
}
