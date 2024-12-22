package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const (
	addr = ":12345"
	// Протокол сетевой службы
	proto = "tcp4"
)

func main() {
	// Создаем клиента слуэбы RPC
	client, err := rpc.DialHTTP(proto, addr)
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var resp string
	// Удаленный вызов процедуры Sever.Time Должна быть ошибка
	err = client.Call("Server.Time", "unknow message", &resp)
	if err != nil {
		fmt.Println("some error: ", err)
	}
	fmt.Println("time: ", resp)
	// Удаленный вызов процедуры Sever.Time Должет быть ответ
	err = client.Call("Server.Time", "time", &resp)
	if err != nil {
		fmt.Println("some error: ", err)
	}
	fmt.Println("time: ", resp)
}
