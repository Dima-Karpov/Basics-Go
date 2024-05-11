package main

import "log"

func main() {
	var sum int
	for i := 0; i < 50; i++ {
		sum += i
	}

	log.Println(sum)
}
