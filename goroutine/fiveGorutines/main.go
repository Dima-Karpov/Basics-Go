package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	routines := 5
	iterations := 10

	// Добавляем routines в WaitGroup
	wg.Add(routines)
	for i := 1; i <= routines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				fmt.Printf("Routine %d, iteration %d\n", id, j+1)
			}
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}
