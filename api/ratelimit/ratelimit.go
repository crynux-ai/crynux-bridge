package ratelimit

import (
	"context"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var (
	APIRateLimiter = NewRateLimiter()
)

type RateLimiter struct {
	mu       sync.RWMutex
	limiters map[string]*rate.Limiter
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
}

func (r *RateLimiter) CheckRateLimit(ctx context.Context, key string, limit int64, period time.Duration) (bool, float64, error) {
	r.mu.Lock()
	limiter, exists := r.limiters[key]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(limit)*rate.Every(period), int(limit))
		r.limiters[key] = limiter
	}
	r.mu.Unlock()

	reservation := limiter.Reserve()
	if !reservation.OK() {
		return false, 0, nil
	}

	waitTime := reservation.Delay()
	if waitTime > 0 {
		return false, waitTime.Seconds(), nil
	}

	return true, 0, nil
}

func (r *RateLimiter) UpdateRateLimit(ctx context.Context, key string, limit int64, period time.Duration) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	limiter, exists := r.limiters[key]
	if !exists {
		limiter = rate.NewLimiter(rate.Limit(limit)*rate.Every(period), int(limit))
		r.limiters[key] = limiter
	} else {
		limiter.SetLimit(rate.Limit(limit)*rate.Every(period))
	}

	return nil
}
