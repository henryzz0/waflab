package docker

import (
	"sync"
)

type safeCounter struct {
	mux sync.Mutex
	value int
}

func (s *safeCounter) Increment() {
	s.mux.Lock()
	s.value += 1
	s.mux.Unlock()
}

func (s *safeCounter) Decrement() {
	s.mux.Lock()
	s.value -= 1
	s.mux.Unlock()
}

func (s *safeCounter) Value() int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.value
}