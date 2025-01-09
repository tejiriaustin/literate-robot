package middleware

import (
	"context"
	"strings"

	"github.com/tejiriaustin/literate-robot/core/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"user-service/model"
)

type AuthInterceptor struct {
	jwtManager *jwt.JWTManager[model.User]
}

func NewAuthInterceptor(jwtManager *jwt.JWTManager[model.User]) *AuthInterceptor {
	return &AuthInterceptor{
		jwtManager: jwtManager,
	}
}

func (i *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if !requiresAuth(info.FullMethod) {
			return handler(ctx, req)
		}

		userID, err := i.authorize(ctx)
		if err != nil {
			return nil, err
		}

		ctx = context.WithValue(ctx, "user_id", userID)
		return handler(ctx, req)
	}
}

func (i *AuthInterceptor) authorize(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", status.Errorf(codes.Unauthenticated, "metadata is not provided")
	}

	values := md["authorization"]
	if len(values) == 0 {
		return "", status.Errorf(codes.Unauthenticated, "authorization token is not provided")
	}

	accessToken := values[0]
	user, err := i.jwtManager.Validate(accessToken)
	if err != nil {
		return "", status.Errorf(codes.Unauthenticated, "access token is invalid: %v", err)
	}

	return user.ID.String(), nil
}

func requiresAuth(method string) bool {
	if strings.Contains(method, "/auth") {
		return false
	}
	return true
}
