package server

import pb "github.com/new-timlieberman/gitasy2.0/proto"

type Server struct {
	authClient pb.AuthServiceClient
}

func New(
	authClient pb.AuthServiceClient,
) *Server {
	return &Server{
		authClient: authClient,
	}
}
