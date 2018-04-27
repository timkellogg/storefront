package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9119"
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
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	RegisterIdentificationServiceServer(grpcServer, &Server{})
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Identification service failed: %v", err)
	}
}
