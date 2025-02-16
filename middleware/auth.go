package middleware

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	"google.golang.org/grpc/metadata"
	"url-shortener/auth"
)

type AuthInterceptor struct{}

func (interceptor *AuthInterceptor) Validate(
	ctx context.Context,
	accessibleMethods map[string]bool,
) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("No metadata found")
		return status.Errorf(codes.Unauthenticated, "missing metadata")
	}

	tokens := md["authorization"]
	if len(tokens) == 0 {
		log.Println("Missing token")
		return status.Errorf(codes.Unauthenticated, "missing token")
	}

	_, err := auth.ValidateJWT(tokens[0])
	if err != nil {
		log.Printf("Invalid token: %v", err)
		return status.Errorf(codes.Unauthenticated, "invalid token")
	}

	return nil
}
