package main

import (
	"log"
	"net"

	auth "github.com/new-timlieberman/gitasy2.0/auth/internal/server"
	pb "github.com/new-timlieberman/gitasy2.0/proto/auth"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	//TODO: create enviornemt variable.
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	conn, err := grpc.NewClient(
		"user:50052",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	userClient := userpb.NewUserServiceClient(conn)

	rdb := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	authServer := auth.New(
		userClient,
		rdb,
		"secret",
	)

	pb.RegisterAuthServiceServer(grpcServer, authServer)

	log.Println("auth service running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
