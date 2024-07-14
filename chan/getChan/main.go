package main

import "fmt"

func getChan() <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 1; i <= 5; i++ {
			c <- i
		}
	}()

	return c
}

func main() {
	c := getChan()
	for i := range c {
		fmt.Printf("%v\n", i)
	}
}
