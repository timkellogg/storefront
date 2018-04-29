package main

import (
	micro "github.com/micro/go-micro"
	userProto "github.com/timkellogg/store/user/protos/user"
)

var config = Config{}

func main() {
	config.Read()

	service := micro.NewService(
		micro.Name("go.micro.api"),
		micro.Version("0.0.1"),
	)

	userProto.RegisterUserServiceHandler(service.Server(), new(API))
}
