package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/new-timlieberman/gitasy2.0/api/internal/middleware"

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
	api.POST("/refresh", authHandler.Refresh)

	protected := api.Group("")
	protected.Use(middleware.JWTMiddleware)
	protected.GET("/users/:id", userHandler.GetUser)
}
