package user

import (
	"github.com/new-timlieberman/gitasy2.0/internal/db"
	pb "github.com/new-timlieberman/gitasy2.0/proto/user"
)

type Server struct {
	pb.UnimplementedUserServiceServer

	queries *db.Queries
}

func New() *Server {
	return &Server{}
}
