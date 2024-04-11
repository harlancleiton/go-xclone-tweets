package grpc

import (
	"context"
	"strings"
	"time"

	"github.com/harlancleiton/go-tweets/internal/domain/entities"
	"github.com/harlancleiton/go-tweets/internal/domain/repositories"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type contextKey string

const (
	AuthenticatedUser contextKey = "authenticated_user"
)

const (
	BearerToken = "Bearer"
)

var (
	ErrUnauthorized = status.Errorf(codes.Unauthenticated, "unauthorized")
	ErrInternal     = status.Errorf(codes.Internal, "unauthorized")
)

type AuthInterceptor struct {
	userRepository repositories.UserRepository
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		user, err := authorize(ctx, interceptor.userRepository)

		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, AuthenticatedUser, user)
		return handler(ctx, req)
	}
}

func authorize(ctx context.Context, userRepository repositories.UserRepository) (*entities.Author, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, ErrUnauthorized
	}

	authorizationHeader := md.Get("authorization")

	if len(authorizationHeader) == 0 {
		return nil, ErrUnauthorized
	}

	accessToken := strings.TrimPrefix(authorizationHeader[0], BearerToken+" ")
	mySigningKey := []byte("secret")
	token, err := jwt.ParseWithClaims(
		accessToken,
		&jwt.StandardClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)

			if !ok {
				return nil, ErrInternal
			}

			return mySigningKey, nil
		},
	)

	if err != nil {
		return nil, ErrUnauthorized
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)

	if !ok {
		return nil, ErrInternal
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	return userRepository.FindByUsername(ctx, claims.Subject)
}

func NewAuthInterceptor(userRepository repositories.UserRepository) *AuthInterceptor {
	return &AuthInterceptor{
		userRepository: userRepository,
	}
}
