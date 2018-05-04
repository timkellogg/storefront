package main

import (
	"context"

	proto "github.com/timkellogg/store/account/protos/account"
)

// Server - represents the gRPC server
type Server struct{}

// GetAccount - find account by id
func (s *Server) GetAccount(ctx context.Context, req *proto.GetAccountRequest, res *proto.Account) error {
	account, err := accountRepository.FindByID(req.Id)
	res.Email = account.Email
	res.Guid = account.Guid
	res.Name = account.Name
	return err
}

// CreateAccount - persist account to the database
func (s *Server) CreateAccount(ctx context.Context, req *proto.Account, res *proto.Account) error {
	// validate
	err := accountRepository.Insert(req)
	return err
}
