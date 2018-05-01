package main

import (
	"log"
	"time"

	micro "github.com/micro/go-micro"
	accountProto "github.com/timkellogg/store/account/protos/account"
)

var accountRepository = AccountRepository{}
var config = Config{}

func init() {
	config.Read()

	accountRepository.Server = config.Server
	accountRepository.Database = config.Database
	accountRepository.Connect()
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.accounts"),
		micro.Version("0.0.1"),
		micro.RegisterTTL(time.Second*30),
	)

	service.Init()

	accountProto.RegisterAccountsHandler(service.Server(), new(Server))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
