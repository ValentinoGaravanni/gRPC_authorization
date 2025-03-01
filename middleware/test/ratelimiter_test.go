package test

import (
	"context"
	"testing"
	"time"
	"url-shortener/middleware"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestRateLimitingInterceptor(t *testing.T) {
	// 3 request in a sec max
	rl := middleware.NewRateLimiter(3, time.Second)
	interceptor := middleware.RateLimitingInterceptor(rl)

	// Imitation of website call
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "OK", nil
	}

	ctx := context.Background()

	//check
	for i := 0; i < 3; i++ {
		_, err := interceptor(ctx, nil, &grpc.UnaryServerInfo{}, handler)
		if err != nil {
			t.Errorf("Expected no error for request %d, got: %v", i+1, err)
		}
	}

	// 4th request
	_, err := interceptor(ctx, nil, &grpc.UnaryServerInfo{}, handler)
	if err == nil || status.Code(err) != codes.ResourceExhausted {
		t.Errorf("Expected ResourceExhausted error, got: %v", err)
	}

	// limit reset
	time.Sleep(time.Second)

	// Test if we can make a request
	_, err = interceptor(ctx, nil, &grpc.UnaryServerInfo{}, handler)
	if err != nil {
		t.Errorf("Expected no error after rate limit reset, got: %v", err)
	}
}
