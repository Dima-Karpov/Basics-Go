package main

import (
	"bufio"
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"strings"
	"time"
)

// Сетевой адрес
//
// Служба будет слушать запрсы на всех IP-адресах
// компьютер на порту 12345
// Например, 127.0.0.1:12345
const (
	addr = "0.0.0.0:12345"
	// Протокол сетевой службы
	proto = "tcp4"
)

// Тип данных для RPC сервера.
// Может быть любым пользовательским типом.
// Все экспортируемые методы этого типа
// будут доступны для удаленного вызова.
type Server int

// Обработчик. Вызывается для каждого соединения
func handleConn(conn net.Conn) {
	// Закрытие соединения
	defer conn.Close()
	// Чтение сообщнеия от клиента
	reader := bufio.NewReader(conn)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		log.Println(err)
		return
	}

	// Удаление символов конца строки
	msg := strings.TrimSuffix(string(b), "\n")
	msg = strings.TrimSuffix(msg, "\r")
	// Если получили "time" - пишем время в соединение
	if msg == "time" {
		conn.Write([]byte("this time" + " " + time.Now().String() + "\n"))
	}
}

// Метод Time доступен для удаленногно вызова
func (s *Server) Time(msg string, resp *string) error {
	if msg != "time" {
		return errors.New("invalid msg")
	}
	*resp = time.Now().String()
	return nil
}

func main() {
	// simple
	//// Запуск сетевой службы по протоколу TCP
	//// на порту 12345
	//listener, err := net.Listen(proto, addr)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//// defer listener.Close()
	//
	//// Подключения обрабатываются в бесконечном цикле.
	//// Иначе после обслуживания первого подключения сервер
	//// завершит работу
	//for {
	//	// Принимаем подключения
	//	conn, err := listener.Accept()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	// Вызов обработчика подключения
	//	go handleConn(conn)
	//}

	// RPC
	// Создаем указатель на переменную типа Server
	srv := new(Server)
	// Регистрируем методы типа Server в службе RPC
	rpc.Register(srv)
	// Регистрируем HTTP-обработчик для службы RPC
	rpc.HandleHTTP()
	// Создаем сетевую службу
	listener, err := net.Listen(proto, addr)
	if err != nil {
		log.Fatal(err)
	}
	// Запускаем HTTP-сервер поверх созданной сетевой службы
	http.Serve(listener, nil)
}
