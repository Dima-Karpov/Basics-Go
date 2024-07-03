package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {

	// Считываем значение n из консоли
	fmt.Print("Введите количество горутин: ")
	var n int

	fmt.Scanln(&n)

	// Канал для уведомления о завершении программы
	done := make(chan bool)

	// Запускаем n goroutine
	for i := 1; i <= n; i++ {
		go func(id int) {
			for {
				select {
				case <-done:
					return
				default:
					fmt.Printf("Горутина %d\n", id)
					time.Sleep(1 * time.Second)
				}
			}
		}(i)
	}

	// Ожидаем нажатия Enter
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// Уведомляем все горутины о завершении
	close(done)

	// Немного подождем, чтобы все горутины завершились корректно
	time.Sleep(1 * time.Second)
	fmt.Println("Программа завершена")

}
