package main

import (
	"log"
	"net"

	identificationProto "github.com/timkellogg/store/identification/protos/identification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":8080"
)

var identityRepository = IdentityRepository{}
var config = Config{}

func init() {
	config.Read()

	identityRepository.Server = config.Server
	identityRepository.Database = config.Database
	identityRepository.Connect()
}

func main() {
	log.Println("Starting identification service...")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	identificationProto.RegisterIdentificationServiceServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Identification service failed to start: %v", err)
	}
}
