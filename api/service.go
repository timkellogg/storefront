package main

import (
	"log"
	"time"

	web "github.com/micro/go-web"
)

var config = Config{}

func main() {
	config.Read()

	// service := micro.NewService(
	// 	micro.Name("go.micro.api"),
	// 	micro.Version("0.0.1"),
	// )

	// accountProto.RegisterAccountServiceHandler(service.Server(), new(API))

	service := web.NewService(
		web.Name("go.micro.api"),
		web.Version("0.0.1"),
		web.RegisterTTL(time.Second*30),
	)

	service.Init()

	service.HandleFunc("/accounts/{accountID}", getAccountHandler)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
