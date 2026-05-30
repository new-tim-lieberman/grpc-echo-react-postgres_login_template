package routes

import (
	"github.com/labstack/echo/v4"

	authpb "github.com/new-timlieberman/gitasy2.0/proto/auth"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"

	"github.com/new-timlieberman/gitasy2.0/api/internal/handlers"
)

func RegisterRoutes(
	e *echo.Echo,
	authClient authpb.AuthServiceClient,
	userClient userpb.UserServiceClient,
) {
	authHandler := handlers.NewAuthHandler(authClient)
	userHandler := handlers.NewUserHandler(userClient)

	//e.GET("/health", handlers.Health)

	api := e.Group("/api")

	api.POST("/login", authHandler.Login)
	api.POST("/register", authHandler.Register)

	api.GET("/users/:id", userHandler.GetUser)
}
