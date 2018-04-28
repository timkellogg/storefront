package main

import (
	"log"
	"time"

	micro "github.com/micro/go-micro"
	proto "github.com/timkellogg/store/user/protos/user"
)

var userRepository = UserRepository{}
var config = Config{}

func init() {
	config.Read()

	userRepository.Server = config.Server
	userRepository.Database = config.Database
	userRepository.Connect()
}

func main() {
	service := micro.NewService(
		micro.Name("user.service"),
		micro.RegisterTTL(time.Second*30),
	)

	service.Init()

	proto.RegisterUsersServiceHandler(service.Server(), new(Server))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
