package gorputil

import (
	"math/rand"
	"sync"
)

type Balancer interface {
	Balance(n int) int
}

type Sequential struct {
	mu      sync.Mutex
	counter int
}

func (s *Sequential) Balance(n int) int {
	s.mu.Lock()
	defer s.mu.Unlock()
	r := s.counter
	s.counter = (s.counter + 1) % n
	return r
}

type Random struct{}

func (r *Random) Balance(n int) int {
	return rand.Intn(n)
}
