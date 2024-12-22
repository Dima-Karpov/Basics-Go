package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strings"
	"sync"
	"time"
)

// Сетевой адрес
const (
	addr  = "0.0.0.0:12345"
	proto = "tcp"
)

// Список Go-поговорок
var proverbs = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}

// Обработчик соединений
func handleConn(conn net.Conn) {
	defer conn.Close()

	stop := make(chan bool)
	var wg sync.WaitGroup

	// Горутина для отправки поговорок
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				return
			default:
				// Выбираем случайную поговорку
				proverb := proverbs[rand.Intn(len(proverbs))]
				// Отправляем клиенту
				_, err := conn.Write([]byte(proverb + "\n"))
				if err != nil {
					fmt.Println("Ошибка отправки данных клиенту:", err)
					return
				}
				// Ждем 3 секунды
				time.Sleep(3 * time.Second)
			}
		}
	}()

	// Чтения команд от клиента
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения данных шт клиента:", err)
			close(stop)
			break
		}

		command := strings.TrimSpace(message)
		if command == "stop" {
			fmt.Println("Получена команда stop от", conn.RemoteAddr())
			close(stop)
			break
		}
	}

	// Ожидаем завершения горутины
	wg.Wait()
}

func main() {
	listener, err := net.Listen(proto, addr)
	if err != nil {
		fmt.Println("Не удалось запустить сервер:", err)
		return
	}
	defer listener.Close()
	fmt.Println("Сервер запущен на", addr)

	for {
		// Принимем новое подключение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Ошибка подключения:", err)
			continue
		}
		fmt.Println("Новое подключение от", conn.RemoteAddr())
		go handleConn(conn)
	}
}
