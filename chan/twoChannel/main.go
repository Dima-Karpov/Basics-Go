package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Создаем два канала для сообщений типа int
	channel1 := make(chan int)
	channel2 := make(chan int)

	// Запускаем горутину для отправки сообщений в channel1
	go func() {
		for i := 0; ; i++ {
			channel1 <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	// Запускаем горутину для отправки сообщений во channel2
	go func() {
		for i := 0; ; i++ {
			channel2 <- i
			time.Sleep(time.Millisecond * 100)
		}
	}()

	for {
		select {
		// Приём сообщения из channel1 с вероятностью 2/3
		case msg := <-channel1:
			if rand.Float64() < 2.0/3.0 {
				fmt.Printf("Received from channel1: %d\n", msg)
			} else {
				// Если условие не выполняется, продолжаем без блокировки
				continue
			}
		// Приём сообщения из channel2 с вероятностью 1/3
		case msg := <-channel2:
			if rand.Float64() < 1.0/3.0 {
				fmt.Printf("Received from channel2: %d\n", msg)
			} else {
				// Если условие не выполняется, продолжаем без блокировки
				continue
			}
		}
	}
}
