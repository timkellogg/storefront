package main

import (
	"log"

	context "golang.org/x/net/context"
)

// Server - represents the gRPC server
type Server struct{}

// Get - find identity by id
func (s *Server) Get(ctx context.Context, in *GetRequest) (*Identity, error) {
	identity, err := identityRepository.FindByID(in.Id)
	if err != nil {
		log.Println(err)
	}
	return &identity, err
}
