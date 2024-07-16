package main

import (
	"fmt"
	"sync"
)

// Функция отправитель, которая отправляет данные в канал и закрывает его
func sender(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// Функция получатель, которая принимает данные из канала
func receiver(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println("Received: ", num)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)

	// Запуск горутины отправителя
	go sender(ch, &wg)

	// Запуск горутины получателя
	go receiver(ch, &wg)

	// Ожидание завершения всех горутин
	wg.Wait()
}
