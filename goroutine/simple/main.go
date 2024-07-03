package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	//for i := 0; i < 3; i++ {
	//	go func() {
	//		fmt.Println("Hello, goroutine!")
	//		time.Sleep(time.Second)
	//	}()
	//}
	//time.Sleep(time.Second)

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	fmt.Println(<-c + <-c)
}
