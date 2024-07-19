package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(from int, to int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем  счетчик на 1, когда функция завершится
	for i := from; i <= to; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go printNumbers(1, 5, &wg)
	go printNumbers(6, 10, &wg)

	wg.Wait()
	fmt.Println("done")
}
