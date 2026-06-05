package user

import (
	"context"
	"testing"

	"github.com/new-timlieberman/gitasy2.0/internal/db"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
)

func TestGetUserByEmail(t *testing.T) {
	mockQueries := &MockQueries{
		GetUserByEmailFunc: func(
			ctx context.Context,
			email string,
		) (db.User, error) {

			return db.User{
				ID:           1,
				Email:        email,
				PasswordHash: "hashed_password",
			}, nil
		},
	}

	server := New(mockQueries)

	resp, err := server.GetUserByEmail(
		context.Background(),
		&userpb.GetUserByEmailRequest{
			Email: "test@example.com",
		},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Email != "test@example.com" {
		t.Fatalf(
			"expected email test@example.com, got %s",
			resp.Email,
		)
	}

	if resp.Id != 1 {
		t.Fatalf(
			"expected id 1, got %d",
			resp.Id,
		)
	}
}
