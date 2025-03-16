package main

import "fmt"

func Fibonacci(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	fmt.Println("F(10):", Fibonacci(10)) // 55
	fmt.Println("F(5):", Fibonacci(5))   // 5
	fmt.Println("F(1):", Fibonacci(1))   // 1
}
