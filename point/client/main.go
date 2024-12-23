package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const addr = ":12345"
const network = "tcp4"

// Точка на плоскости.
type Point struct {
	X, Y float64
}

// Аргумент для функции Dist.
type Points struct {
	A, B Point
}

func main() {
	// Создаем клиента службы RPC.
	client, err := rpc.DialHTTP(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	var p Points = Points{
		A: Point{X: 1, Y: 1},
		B: Point{X: 4, Y: 5},
	}
	var resp float64
	err = client.Call("Server.Dist", p, &resp)
	if err != nil {
		fmt.Println("ошибка:", err)
	}
	fmt.Println("dist:", resp)
}
