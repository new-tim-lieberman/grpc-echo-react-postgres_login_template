package server

import (
	"context"
	"time"

	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"

	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
)

type MockUserClient struct {
	CreateUserFunc func(
		ctx context.Context,
		in *userpb.CreateUserRequest,
		opts ...grpc.CallOption,
	) (*userpb.UserResponse, error)

	GetUserByEmailFunc func(
		ctx context.Context,
		in *userpb.GetUserByEmailRequest,
		opts ...grpc.CallOption,
	) (*userpb.UserResponse, error)
}

func (m *MockUserClient) CreateUser(
	ctx context.Context,
	in *userpb.CreateUserRequest,
	opts ...grpc.CallOption,
) (*userpb.UserResponse, error) {
	return m.CreateUserFunc(ctx, in, opts...)
}

func (m *MockUserClient) GetUserByEmail(
	ctx context.Context,
	in *userpb.GetUserByEmailRequest,
	opts ...grpc.CallOption,
) (*userpb.UserResponse, error) {
	return m.GetUserByEmailFunc(ctx, in, opts...)
}

type MockRedis struct {
	SetFunc func(
		ctx context.Context,
		key string,
		value interface{},
		expiration time.Duration,
	) *redis.StatusCmd

	GetFunc func(
		ctx context.Context,
		key string,
	) *redis.StringCmd
}

func (m *MockRedis) Set(
	ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration,
) *redis.StatusCmd {
	return m.SetFunc(ctx, key, value, expiration)
}

func (m *MockRedis) Get(
	ctx context.Context,
	key string,
) *redis.StringCmd {
	return m.GetFunc(ctx, key)
}
