package server

import (
	"github.com/new-timlieberman/gitasy2.0/internal/db"
	pb "github.com/new-timlieberman/gitasy2.0/proto"
)

type Server struct {
	pb.UnimplementedAuthServiceServer

	queries *db.Queries
}
