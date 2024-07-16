package main

import (
	"fmt"
	"sync"
)

// Counter is safe to use concurrently via channels
type Counter struct {
	val    int
	incr   chan int
	getVal chan chan int
}

// NewCounter creates a new Counter
func NewCounter() *Counter {
	c := &Counter{
		incr:   make(chan int),
		getVal: make(chan chan int),
	}

	go c.run()
	return c
}

// run is the goroutine that listens on channels and updates the counter
func (c *Counter) run() {
	for {
		select {
		case delta := <-c.incr:
			c.val += delta
		case ch := <-c.getVal:
			ch <- c.val
		}
	}
}

// Inc increments the counter by a given value
func (c *Counter) Inc(delta int) {
	c.incr <- delta
}

// Value returns the current value of the counter
func (c *Counter) Value() int {
	ch := make(chan int)
	c.getVal <- ch
	return <-ch
}

func main() {
	c := NewCounter()
	var wg sync.WaitGroup

	// Start 1000 goroutines to increment the counter
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			c.Inc(1)
			wg.Done()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
	fmt.Println(c.Value())
}
