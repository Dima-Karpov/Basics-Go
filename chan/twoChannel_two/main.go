package main

import (
	"fmt"
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
			time.Sleep(2 * time.Second) // каждые 2 секунды
		}
	}()

	// Запускаем горутину для отправки сообщений в channel2
	go func() {
		for i := 0; ; i++ {
			channel1 <- i * 10
			time.Sleep(3 * time.Second) // каждые 3 секунды
		}
	}()

	// Главная горутина для приема сообщений из обоих каналов
	for {
		select {
		default:
			fmt.Println("Текущее время: ", time.Now().Format("15:04:05"))
			time.Sleep(1 * time.Second) // пауза 1 секунда

		case msg := <-channel1:
			fmt.Println("Получено сообщение из channel1: ", msg)
		case msg := <-channel2:
			fmt.Println("Получено сообщение из channel2: ", msg)
		}
	}
}
