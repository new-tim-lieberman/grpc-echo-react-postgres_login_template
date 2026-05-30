package user

import (
	"context"

	"github.com/new-timlieberman/gitasy2.0/internal/db"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
)

type Server struct {
	userpb.UnimplementedUserServiceServer
	queries *db.Queries
}

func New() *Server {
	return &Server{}
}

func (s *Server) GetUser(
	ctx context.Context,
	req *userpb.GetUserRequest,
) (*userpb.UserResponse, error) {

	return &userpb.UserResponse{
		Id: req.Id,
	}, nil
}

func (s *Server) CreateUser(
	ctx context.Context,
	req *userpb.CreateUserRequest,
) (*userpb.UserResponse, error) {
	user, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Email: req.Email,
	})

	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{
		Id:    int32(user.ID),
		Email: user.Email,
	}, nil
}
