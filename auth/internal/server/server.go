package server

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	authpb "github.com/new-timlieberman/gitasy2.0/proto/auth"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer

	userClient UserClient
	rdb        RedisClient
	jwtSecret  string
}

func New(
	userClient UserClient,
	rdb RedisClient,
	jwtSecret string,
) *Server {
	return &Server{
		userClient: userClient,
		rdb:        rdb,
		jwtSecret:  jwtSecret,
	}
}

type UserClient interface {
	CreateUser(
		ctx context.Context,
		in *userpb.CreateUserRequest,
		opts ...grpc.CallOption,
	) (*userpb.UserResponse, error)

	GetUserByEmail(
		ctx context.Context,
		in *userpb.GetUserByEmailRequest,
		opts ...grpc.CallOption,
	) (*userpb.UserResponse, error)
}

type RedisClient interface {
	Set(
		ctx context.Context,
		key string,
		value interface{},
		expiration time.Duration,
	) *redis.StatusCmd

	Get(
		ctx context.Context,
		key string,
	) *redis.StringCmd
}

func (s *Server) Register(
	ctx context.Context,
	req *authpb.RegisterRequest,
) (*authpb.RegisterResponse, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return nil, err
	}
	fmt.Println("hashedPassword:", string(hashedPassword))
	fmt.Println(string(hashedPassword))
	userResponse, err := s.userClient.CreateUser(
		ctx,
		&userpb.CreateUserRequest{
			Email:        req.Email,
			PasswordHash: string(hashedPassword),
		},
	)
	if err != nil {
		return nil, err
	}

	accessToken, refreshToken, err := s.createTokens(
		ctx,
		userResponse.Id,
	)
	if err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{
		Message:      "registered",
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) Login(
	ctx context.Context,
	req *authpb.LoginRequest,
) (*authpb.LoginResponse, error) {

	user, err := s.userClient.GetUserByEmail(
		ctx,
		&userpb.GetUserByEmailRequest{
			Email: req.Email,
		},
	)

	fmt.Println("user:", user)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(req.Password),
	)
	if err != nil {
		return nil, status.Error(
			codes.Unauthenticated,
			"invalid credentials",
		)
	}

	accessToken, refreshToken, err := s.createTokens(
		ctx,
		user.Id,
	)
	if err != nil {
		return nil, err
	}

	return &authpb.LoginResponse{
		Token:        accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Server) createTokens(
	ctx context.Context,
	userID int32,
) (string, string, error) {

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userID,
			"exp": time.Now().
				Add(time.Minute * 15).
				Unix(),
		},
	)

	accessToken, err := token.SignedString(
		[]byte(s.jwtSecret),
	)
	if err != nil {
		return "", "", err
	}

	b := make([]byte, 32)

	_, err = rand.Read(b)
	if err != nil {
		return "", "", err
	}

	refreshToken := hex.EncodeToString(b)

	err = s.rdb.Set(
		ctx,
		refreshToken,
		userID,
		time.Hour*24*30,
	).Err()

	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *Server) RefreshToken(
	ctx context.Context,
	req *authpb.RefreshTokenRequest,
) (*authpb.AuthResponse, error) {

	userID, err := s.rdb.Get(
		ctx,
		req.RefreshToken,
	).Result()

	if err != nil {
		return nil, status.Error(
			codes.Unauthenticated,
			"invalid refresh token",
		)
	}

	id, err := strconv.ParseInt(
		userID,
		10,
		32,
	)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": int32(id),
			"exp": time.Now().
				Add(time.Minute * 15).
				Unix(),
		},
	)

	tokenString, err := token.SignedString(
		[]byte(s.jwtSecret),
	)
	if err != nil {
		return nil, err
	}

	return &authpb.AuthResponse{
		Message: "token refreshed",
		Token:   tokenString,
	}, nil
}
