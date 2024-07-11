package main

import (
	"fmt"
	"sync"
)

// Шаг наращивания счетчика
const step int64 = 1

// Конечное значение счетчика
const interationAmount int = 1000

func main() {
	var counter int64 = 0
	var mu sync.Mutex
	var c = sync.NewCond(&mu)
	done := false

	increment := func(i int) {
		mu.Lock()
		defer mu.Unlock()
		counter += step
		if i == interationAmount {
			done = true
			c.Signal()
		}
	}

	for i := 1; i <= interationAmount; i++ {
		go increment(i)

	}

	mu.Lock()
	for !done {
		c.Wait()
	}
	mu.Unlock()

	fmt.Println("Result:", counter)
}
