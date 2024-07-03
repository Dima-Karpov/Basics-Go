package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

func incrementCounter() {
	for {
		time.Sleep(1 * time.Second)
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func printCounter() {
	for {
		time.Sleep(200 * time.Millisecond)
		mu.Lock()
		fmt.Printf("%v - %v\n", time.Now().Format("15:04:05.00"), counter)
		mu.Unlock()
	}
}

func main() {
	var n int
	fmt.Print("Введите количество секунд для выполнения программы: ")
	fmt.Fscan(os.Stdin, &n)

	go incrementCounter()
	go printCounter()

	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println("Программа завершена.")
}
