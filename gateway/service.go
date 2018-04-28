package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var config = Config{}

func main() {
	log.Println("Establishing gateway service...")
	config.Read()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/identity/{id}", getIdentityHandler).Methods("GET")
	r.HandleFunc("/api/v1/identity", createIdentityHandler).Methods("POST")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
