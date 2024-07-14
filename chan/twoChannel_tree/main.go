package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := make(chan int)
	var wg sync.WaitGroup

	// Первая горутина отправляет числа от 1 до 100 в канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	// Вторая горутина принимает числа из канала и печатает их
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range numbers {
			fmt.Println(num)
		}
	}()

	wg.Wait()
}
