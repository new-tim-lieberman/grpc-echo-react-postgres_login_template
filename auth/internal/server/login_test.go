package server

import (
	"context"
	"testing"
	"time"

	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
	"google.golang.org/grpc"

	authpb "github.com/new-timlieberman/gitasy2.0/proto/auth"
	"github.com/redis/go-redis/v9"

	"golang.org/x/crypto/bcrypt"
)

func TestLogin(t *testing.T) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte("password123"),
		bcrypt.DefaultCost,
	)

	if err != nil {
		t.Fatal(err)
	}

	mockUserClient := &MockUserClient{
		GetUserByEmailFunc: func(
			ctx context.Context,
			req *userpb.GetUserByEmailRequest,
			opts ...grpc.CallOption,
		) (*userpb.UserResponse, error) {

			return &userpb.UserResponse{
				Id:           1,
				Email:        req.Email,
				PasswordHash: string(hashedPassword),
			}, nil
		},
	}

	mockRedis := &MockRedis{
		SetFunc: func(
			ctx context.Context,
			key string,
			value interface{},
			expiration time.Duration,
		) *redis.StatusCmd {

			return redis.NewStatusCmd(ctx)
		},
	}

	server := New(
		mockUserClient,
		mockRedis,
		"super-secret",
	)

	resp, err := server.Login(
		context.Background(),
		&authpb.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		},
	)

	if err != nil {
		t.Fatalf(
			"expected no error, got %v",
			err,
		)
	}

	if resp.Token == "" {
		t.Fatal("expected token")
	}

	if resp.RefreshToken == "" {
		t.Fatal("expected refresh token")
	}
}
