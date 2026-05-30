package main

import (
	"log"

	echo "github.com/labstack/echo/v4"
	"github.com/new-timlieberman/gitasy2.0/api/internal/routes"
	authpb "github.com/new-timlieberman/gitasy2.0/proto/auth"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// -----------------------------
	// gRPC connections
	// -----------------------------

	authConn, err := grpc.NewClient(
		"auth:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to auth service: %v", err)
	}
	defer authConn.Close()

	userConn, err := grpc.NewClient(
		"user:50052",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("failed to connect to user service: %v", err)
	}
	defer userConn.Close()

	authClient := authpb.NewAuthServiceClient(authConn)
	userClient := userpb.NewUserServiceClient(userConn)

	// -----------------------------
	// Echo server
	// -----------------------------
	e := echo.New()
	routes.RegisterRoutes(e, authClient, userClient)
	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"status": "ok",
		})
	})

	log.Println("api service running on :8080")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}

}
