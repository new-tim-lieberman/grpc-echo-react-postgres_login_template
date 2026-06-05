package user

import (
	"context"

	"github.com/new-timlieberman/gitasy2.0/internal/db"
)

type MockQueries struct {
	GetUserByEmailFunc func(
		ctx context.Context,
		email string,
	) (db.User, error)

	CreateUserFunc func(
		ctx context.Context,
		arg db.CreateUserParams,
	) (db.User, error)
}

func (m *MockQueries) GetUserByEmail(
	ctx context.Context,
	email string,
) (db.User, error) {
	return m.GetUserByEmailFunc(ctx, email)
}

func (m *MockQueries) CreateUser(
	ctx context.Context,
	arg db.CreateUserParams,
) (db.User, error) {
	return m.CreateUserFunc(ctx, arg)
}
