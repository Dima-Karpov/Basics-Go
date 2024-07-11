package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Шаг наращивания счетчика
const step int64 = 1

// Конечное значение счетчика
const interationAmount int = 1000

func main() {
	var counter int64 = 0
	var c = sync.NewCond(&sync.Mutex{})
	increment := func(i int) {
		atomic.AddInt64(&counter, step)
		if i == interationAmount {
			c.Signal()
		}
	}

	for i := 1; i <= interationAmount; i++ {
		go increment(i)
	}

	c.L.Lock()
	c.Wait()
	c.L.Unlock()

	fmt.Println("Result:", counter)
}
