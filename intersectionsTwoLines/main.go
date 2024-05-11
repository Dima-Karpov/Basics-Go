package main

import "fmt"

func main() {
	var x, y float64
	var a1, a2 float64 = 3, 4
	var b1, b2 float64 = 1, 4

	if a1 == a2 {
		if b1 == b2 {
			fmt.Println("line 1 equals line 2")
			return
		}
		fmt.Println("not crossing")
		return
	}

	x = (b2 - b1) / (a1 - a2)
	y = a1*x + b1
	fmt.Printf("x=%f, y=%f\n", x, y)
}
