package grpcutil

import (
	"sync"
	"time"
	"errors"
)

var ErrCircuitOpen = errors.New("circuit breaker is open")

type CircuitBreaker struct {
	mu           sync.Mutex
	failures     int
	threshold    int
	lastFailure  time.Time
	timeout      time.Duration
}

func NewBreaker(threshold int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{threshold: threshold, timeout: timeout}
}

func (b *CircuitBreaker) Execute(fn func() error) error {
	b.mu.Lock()
	if b.failures >= b.threshold && time.Since(b.lastFailure) < b.timeout {
		b.mu.Unlock()
		return ErrCircuitOpen
	}
	b.mu.Unlock()

	err := fn()
	b.mu.Lock()
	defer b.mu.Unlock()

	if err != nil {
		b.failures++
		b.lastFailure = time.Now()
		return err
	}

	b.failures = 0
	return nil
}
