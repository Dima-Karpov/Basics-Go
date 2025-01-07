package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

// handleConnection – обрабатывает каждое подключение
func handleConnection(conn net.Conn) {
	defer conn.Close()
	log.Printf("🔗 Новое соединение с %s", conn.RemoteAddr())

	reader := bufio.NewReader(conn)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Printf("❌ Ошибка чтения данных: %v", err)
			}
			break
		}
		data = strings.TrimSpace(data)
		log.Printf("📥 Получено: %s", data)

		// Модифицируем данные: делаем их заглавными
		modifiedData := strings.ToUpper(data) + "\n"

		_, err = conn.Write([]byte("💬 Echo: " + modifiedData))
		if err != nil {
			log.Printf("❌ Ошибка отправки данных: %v", err)
			break
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("❌ Ошибка запуска TCP-сервера: %v", err)
	}
	defer ln.Close()
	log.Println("🔧 TCP-сервер запущен на :8081")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("❌ Ошибка принятия соединения: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
