package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter type
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// ISafeCounter interface
type ISafeCounter interface {
	Inc(key string)
	Value(key string) int
}

// Inc increments safe counter
func (s *SafeCounter) Inc(key string) {
	s.mux.Lock()
	s.v[key]++
	s.mux.Unlock()
}

// Value gets value of counter
func (s *SafeCounter) Value(key string) int {
	s.mux.Lock()
	defer s.mux.Unlock()
	return s.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)

	fmt.Println(c.Value("somekey"))
}
