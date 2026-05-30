package server

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/new-timlieberman/gitasy2.0/internal/db"
	authpb "github.com/new-timlieberman/gitasy2.0/proto/auth"
	userpb "github.com/new-timlieberman/gitasy2.0/proto/user"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type Server struct {
	authpb.UnimplementedAuthServiceServer
	queries    *db.Queries
	userClient userpb.UserServiceClient
	rdb        *redis.Client
	jwtSecret  string
}

func New(
	userClient userpb.UserServiceClient,
	rdb *redis.Client,
	jwtSecret string,
) *Server {
	return &Server{
		userClient: userClient,
		rdb:        rdb,
		jwtSecret:  jwtSecret,
	}
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

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": userResponse.Id,
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

	b := make([]byte, 32)

	_, err = rand.Read(b)
	if err != nil {
		return nil, err
	}

	refreshToken := hex.EncodeToString(b)

	err = s.rdb.Set(
		ctx,
		refreshToken,
		userResponse.Id,
		time.Hour*24*30,
	).Err()

	if err != nil {
		return nil, err
	}

	return &authpb.RegisterResponse{
		Message:      "registered",
		Token:        tokenString,
		RefreshToken: refreshToken,
	}, nil
}
