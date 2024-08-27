package main

import (
	"fmt"
	"strconv"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func toString(in <-chan int) <-chan string {
	out := make(chan string)
	go func() {
		for n := range in {
			out <- strconv.Itoa(n)
		}
		close(out)
	}()
	return out
}

func main() {
	gen := generator(2, 3, 4)
	sq := square(gen)
	str := toString(sq)

	for s := range str {
		fmt.Println(s)
	}
}
