package main

import (
	"context"

	proto "github.com/timkellogg/store/user/protos/user"
)

// Server - represents the gRPC server
type Server struct{}

// Get - find user by id
func (s *Server) Get(ctx context.Context, req *proto.GetRequest, i *proto.User) error {
	user, err := userRepository.FindByID(req.Id)
	i.Email = user.Email
	i.Guid = user.Guid
	i.Name = user.Name
	return err
}
