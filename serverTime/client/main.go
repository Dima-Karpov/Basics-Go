package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	addr = "localhost:12345"
	// Протокол сетевой службы
	proto = "tcp4"
)

func main() {
	// Подключения к серверу
	conn, err := net.Dial(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Запись запроса
	_, err = conn.Write([]byte("time\n"))
	if err != nil {
		log.Fatal(err)
	}

	// Буфер для чтения данных из соединения
	reader := bufio.NewReader(conn)
	// Считывание массива байт до перевода строки
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Обратока ответа
	fmt.Println("Ответ от сервера:", string(b))
}
