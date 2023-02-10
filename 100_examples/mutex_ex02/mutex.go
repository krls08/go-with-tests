package main

import (
	"fmt"
	"sync"
)

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{

		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for i := 0; i < n; i++ {
			c.inc(name)
			c.mu.Lock()
			fmt.Println("cc ->", c.counters[name])
			c.mu.Unlock()
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10)
	go doIncrement("a", 10)
	go doIncrement("b", 10)

	wg.Wait()
	fmt.Println(c.counters)
}
