package user

import (
	"context"
	"log"

	"github.com/new-timlieberman/gitasy2.0/internal/db"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
)

type Querier interface {
	GetUserByEmail(ctx context.Context, email string) (db.User, error)

	CreateUser(
		ctx context.Context,
		arg db.CreateUserParams,
	) (db.User, error)
}

type Server struct {
	userpb.UnimplementedUserServiceServer
	queries Querier
}

func New(queries Querier) *Server {
	return &Server{
		queries: queries,
	}
}

func (s *Server) GetUserByEmail(
	ctx context.Context,
	req *userpb.GetUserByEmailRequest,
) (*userpb.UserResponse, error) {

	user, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		log.Printf("GetUserByEmail: %v", err)
		return nil, err
	}

	return &userpb.UserResponse{
		Id:           int32(user.ID),
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
	}, nil
}

func (s *Server) CreateUser(
	ctx context.Context,
	req *userpb.CreateUserRequest,
) (*userpb.UserResponse, error) {

	user, err := s.queries.CreateUser(
		ctx,
		db.CreateUserParams{
			Email:        req.Email,
			PasswordHash: req.PasswordHash,
		},
	)

	if err != nil {
		return nil, err
	}

	return &userpb.UserResponse{
		Id:    int32(user.ID),
		Email: user.Email,
	}, nil
}
