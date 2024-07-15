package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value     int
	threshold int
	mu        sync.Mutex
	done      chan struct{}
}

// NewCounter - конструктор для создания нового счетчика
func NewCounter(threshold int) *Counter {
	return &Counter{
		value:     0,
		threshold: threshold,
		done:      make(chan struct{}),
	}
}

// Increment - метод для увеличения значения счетчика
func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.value < c.threshold {
		c.value++
		fmt.Printf("Counter value: %d\n", c.value)
		if c.value >= c.threshold {
			close(c.done)
		}
	}
}

// Done - метод для получения канала завершения
func (c *Counter) Done() <-chan struct{} {
	return c.done
}

func main() {
	var numGoroutines int
	var threshold int

	fmt.Print("Enter the number of goroutines: ")
	fmt.Scan(&numGoroutines)
	fmt.Print("Enter the threshold value: ")
	fmt.Scan(&threshold)

	counter := NewCounter(threshold)
	var wg sync.WaitGroup

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1) // Увеличиваем счетчик WaitGroup
		go func() {
			defer wg.Done() // Уменьшаем счетчик WaitGroup по завершению работы горутины
			for {
				select {
				case <-counter.Done():
					return
				default:
					counter.Increment()
				}
			}
		}()
	}
	wg.Wait()
	fmt.Println("Counter reached the threshold value: ", threshold)
}
