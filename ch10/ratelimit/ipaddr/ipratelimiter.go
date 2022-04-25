package main

import (
	"sync"

	"golang.org/x/time/rate"
)

type ipRateLimiter struct {
	ips map[string]*rate.Limiter
	mu  sync.Mutex
	r   rate.Limit
	b   int
}

func NewIpRateLimiter(r rate.Limit, b int) *ipRateLimiter {
	return &ipRateLimiter{
		ips: make(map[string]*rate.Limiter),
		r:   r,
		b:   b,
	}
}

func (i *ipRateLimiter) Get(ip string) *rate.Limiter {
	i.mu.Lock()
	defer i.mu.Unlock()
	limiter, exists := i.ips[ip]
	if !exists {
		l := rate.NewLimiter(i.r, i.b)
		i.ips[ip] = l
		return l
	}
	return limiter
}
