package main

import (
	"context"

	proto "github.com/timkellogg/store/user/protos/user"
)

// Server - represents the gRPC server
type Server struct{}

// Get - find user by id
func (s *Server) Get(ctx context.Context, req *proto.GetUserRequest, res *proto.User) error {
	user, err := userRepository.FindByID(req.Id)
	res.Email = user.Email
	res.Guid = user.Guid
	res.Name = user.Name
	return err
}
