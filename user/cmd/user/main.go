package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/new-timlieberman/gitasy2.0/internal/db"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
	user "github.com/new-timlieberman/gitasy2.0/user/internal/server"

	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

func main() {
	//TODO: create enviornemt variable.
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	conn, err := sql.Open(
		"postgres",
		"postgres://postgres:postgres@postgres:5432/gitasy?sslmode=disable",
	)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	queries := db.New(conn)

	grpcServer := grpc.NewServer()
	userServer := user.New(queries)

	userpb.RegisterUserServiceServer(
		grpcServer,
		userServer,
	)

	log.Println("user service running on :50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
