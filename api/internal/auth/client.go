package auth

import (
	"log"

	pb "github.com/new-timlieberman/gitasy2.0/proto"
)

func NewClient(addr string) pb.AuthServiceClient {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	return pb.NewAuthServiceClient(conn)
}
