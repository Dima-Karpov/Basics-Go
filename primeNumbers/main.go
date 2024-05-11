package main

import (
	"fmt"
	"math"
)

func main() {
	counter := 1
	for i := 3; counter <= 20; i++ {
		c := false
		for j := 2; j < int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				c = true
			}
			if c {
				break
			}
		}
		if !c {
			fmt.Println(i)
			counter++
		}
	}
}
