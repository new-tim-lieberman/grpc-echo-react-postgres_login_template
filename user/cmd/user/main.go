package main

import (
	"log"
	"net"

	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
	user "github.com/new-timlieberman/gitasy2.0/user/internal/server"
	"google.golang.org/grpc"
)

func main() {
	//TODO: create enviornemt variable.
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userServer := user.New()

	userpb.RegisterUserServiceServer(
		grpcServer,
		userServer,
	)

	userpb.RegisterUserServiceServer(grpcServer, userServer)

	log.Println("user service running on :50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
