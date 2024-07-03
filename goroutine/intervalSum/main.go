package main

import (
	"fmt"
	"sync"
	"time"
)

var mu = sync.Mutex{}

const (
	fiftyThousand      = 50_000
	oneHundredThousand = 100_000
)

func intervalSum(destination *int, start, end int) {
	for i := start; i < end; i++ {
		mu.Lock()
		result := *destination
		result += i
		*destination = result
		mu.Unlock()
	}
}

func main() {
	var result int
	go intervalSum(&result, 0, fiftyThousand)
	go intervalSum(&result, fiftyThousand, oneHundredThousand)
	time.Sleep(time.Second)

	fmt.Println(result)

	otherResult := 0

	for i := 0; i < oneHundredThousand; i++ {
		otherResult += i
	}
	fmt.Println(otherResult)
}
