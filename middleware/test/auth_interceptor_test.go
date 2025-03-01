package test

import (
	"context"
	"os"
	"testing"
	"url-shortener/auth"
	"url-shortener/middleware"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestAuthInterceptor_Validate(t *testing.T) {
	interceptor := &middleware.AuthInterceptor{}

	accessibleMethods := map[string]bool{
		"/pb.URLShortener/URLShort": true,
	}

	validToken, _ := auth.GenerateJWT(os.Getenv("SECRET_KEY"))
	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", validToken))
	err := interceptor.Validate(ctx, accessibleMethods)
	if err != nil {
		t.Errorf("Expected no error for valid token, got: %v", err)
	}

	// Invalid token
	ctx = metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", "invalid-token"))
	err = interceptor.Validate(ctx, accessibleMethods)
	if status.Code(err) != codes.Unauthenticated {
		t.Errorf("Expected Unauthenticated error for invalid token, got: %v", err)
	}

	// No token
	ctx = context.Background()
	err = interceptor.Validate(ctx, accessibleMethods)
	if status.Code(err) != codes.Unauthenticated {
		t.Errorf("Expected Unauthenticated error for missing token, got: %v", err)
	}
}
