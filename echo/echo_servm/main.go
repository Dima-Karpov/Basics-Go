package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	ln, _ := net.Listen("tcp", ":51501") // Устанавливаем прослушивания порта
	for {
		fmt.Println("waiting on port 51501")
		conn, _ := ln.Accept()
		go func(cn net.Conn) {
			defer func() { cn.Close() }()
			for { // Запусткаем цикл подключения
				//  Будем прослушивать все сообщения, разделенные \n
				message, err := bufio.NewReader(cn).ReadString('\n')
				if err != nil { // Клиент разорвал коннект
					return
				}
				cn.Write([]byte(message)) // Отправляем обратно клиенту
			}
		}(conn)
	}

}
