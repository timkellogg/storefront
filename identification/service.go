package identification

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":9119"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	RegisterIdentificationServiceServer(grpcServer, _IdentificationService_serviceDesc)
	reflection.Register(grpcServer)
	grpcServer.Serve(lis)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Identification service failed: %v", err)
	}
}
