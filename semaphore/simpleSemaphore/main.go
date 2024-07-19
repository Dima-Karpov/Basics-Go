package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const numGoroutines = 5
	const maxConcurrent = 2

	// Создаем канал, который будет работать как семафор
	sem := make(chan struct{}, maxConcurrent)

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Запрашиваем доступ к ресурсу
			sem <- struct{}{}

			fmt.Printf("Горутина %d начала работать\n", i)
			time.Sleep(2 * time.Second)
			fmt.Printf("Горутина %d завершила работу\n", i)

			// Освобождаем ресурс
			<-sem
		}(i)
	}

	wg.Wait()
}
