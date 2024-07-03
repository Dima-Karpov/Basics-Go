package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("we start")
	go func() {
		for i := 0; i < 15; i++ {
			time.Sleep(time.Millisecond) // Имитиуем некие вычисления
			fmt.Println("Counter is", i)
		}
	}()

	go func() {
		for i := 0; i < 15; i++ {
			time.Sleep(time.Millisecond) // Имитиуем некие вычисления
			fmt.Println("Other counter is", i)
		}
	}()

	for i := 0; i < 15; i++ {
		fmt.Println("Main counter is", i)
	}

	fmt.Println("we finished")

	time.Sleep(time.Second)
}
