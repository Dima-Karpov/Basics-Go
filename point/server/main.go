package main

import (
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

const addr = ":12345"
const network = "tcp4"

// RPC-сервер.
type Server int

// Точка на плоскости.
type Point struct {
	X, Y float64
}

// Аргумент для функции Dist.
type Points struct {
	A, B Point
}

// Вычисление расстояния между точками.
func (s *Server) Dist(p Points, resp *float64) error {
	*resp = (p.A.X-p.B.X)*(p.A.X-p.B.X) + (p.A.Y-p.B.Y)*(p.A.Y-p.B.Y)
	*resp = math.Sqrt(*resp)
	return nil
}

func main() {
	// Создаем указатель на переменную типа Server.
	srv := new(Server)
	// Регистрируем методы типа Server в службе RPC.
	rpc.Register(srv)
	// Регистрируем HTTP-обработчик для службы RPC.
	rpc.HandleHTTP()
	// Создаём сетевую службу.
	listener, err := net.Listen(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	// Запускаем HTTP-сервер поверх созданной сетевой службы.
	http.Serve(listener, nil)
}
