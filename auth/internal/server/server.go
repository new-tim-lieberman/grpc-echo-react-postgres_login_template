package server

import (
	"context"

	"github.com/new-timlieberman/gitasy2.0/internal/db"
	authpb "github.com/new-timlieberman/gitasy2.0/proto/auth"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer
	queries    *db.Queries
	userClient userpb.UserServiceClient
}

func New(userClient userpb.UserServiceClient) *Server {
	return &Server{
		userClient: userClient,
	}
}

func (s *Server) Register(
	ctx context.Context,
	req *authpb.RegisterRequest,
) (*authpb.RegisterResponse, error) {

	_, err := s.userClient.CreateUser(
		ctx,
		&userpb.CreateUserRequest{
			Email:        req.Email,
			PasswordHash: req.Password,
		},
	)

	if err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{
		Message: "registered",
	}, nil
}
