package main

import (
	"context"
	"log"
	"time"

	identificationProto "github.com/timkellogg/store/identification/protos/identification"
	"google.golang.org/grpc"
)

var config = Config{}

func main() {
	log.Println("Establishing gateway service...")
	config.Read()

	conn, err := grpc.Dial(config.Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	identificationService := identificationProto.NewIdentificationServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := identificationService.Get(ctx, &identificationProto.GetRequest{Id: "1"})
	if err != nil {
		log.Fatalf("identification service failed to get: %v", err)
	}
	log.Println("Gateway registered identification service client")
	log.Print(r)

	// r := mux.NewRouter("/api/user", identificationHandler).Methods("GET")

	// r.HandleFunc("/identifications/", identificationHandler).Methods("GET")

	// if err := http.ListenAndServe(":3000", r); err != nil {
	// 	log.Fatal(err)
	// }
}
