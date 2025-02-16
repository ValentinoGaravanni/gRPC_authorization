package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"sync"
	"time"
)

type RateLimiter struct {
	tokens     int
	maxTokens  int
	mutex      sync.Mutex
	refillRate time.Duration
}

func NewRateLimiter(maxTokens int, refillRate time.Duration) *RateLimiter {
	r1 := &RateLimiter{
		tokens:     maxTokens,
		maxTokens:  maxTokens,
		refillRate: refillRate,
	}
	go r1.refill()
	return r1
}

func (rl *RateLimiter) refill() {
	ticker := time.NewTicker(rl.refillRate)
	defer ticker.Stop()
	for range ticker.C {
		rl.mutex.Lock()
		if rl.tokens < rl.maxTokens {
			rl.tokens++
		}
		rl.mutex.Unlock()
	}
}
func (rl *RateLimiter) Allow() bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()
	if rl.tokens > 0 {
		rl.tokens--
		return true
	}
	return false
}

// RateLimitingInterceptor implements rate limiting for RPC calls.
func RateLimitingInterceptor(rl *RateLimiter) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		if !rl.Allow() {
			return nil, status.Errorf(codes.ResourceExhausted, "Too many requests, try again later")
		}
		return handler(ctx, req)
	}
}
