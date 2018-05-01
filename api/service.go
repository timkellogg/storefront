package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	micro "github.com/micro/go-micro"
)

var config = Config{}

func main() {
	config.Read()

	api := new(API)
	router := mux.NewRouter().StrictSlash(true)

	service := micro.NewService(
		micro.Name("go.micro.api"),
		micro.Version("0.0.1"),
		micro.RegisterTTL(time.Second*30),
	)

	service.Init()

	router.HandleFunc("/accounts/{accountID}", api.getAccountHandler).Methods("GET")

	err := http.ListenAndServe(config.Address, router)
	if err != nil {
		log.Fatal(err)
	}
}
