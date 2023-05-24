package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

// Inc increments the counter value
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.value++
}
func (c *Counter) Value() int {
	return c.value
}

func main() {
	fmt.Println("Counter")
}
