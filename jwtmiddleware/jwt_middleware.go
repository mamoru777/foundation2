package jwtmiddleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

type AuthInterceptor struct{}

func NewAuthInterceptor() *AuthInterceptor {
	return &AuthInterceptor{}
}

func (i *AuthInterceptor) Unary(function ...func()) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		for _, f := range function {
			f()
		}
		return handler(ctx, req)
	}
}

func (i *AuthInterceptor) getTokenFromMetadata(md metadata.MD) string {
	authorization := md.Get("authorization")
	if len(authorization) > 0 {
		// Authorization: Bearer <token>
		parts := strings.Split(authorization[0], " ")
		if len(parts) == 2 {
			return parts[1]
		}
	}
	return ""
}
