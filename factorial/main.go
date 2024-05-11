package main

import "fmt"

func factorialIterative(n int) int {
	res := 1

	for i := 1; i <= n; i++ {
		res *= i
	}

	return res
}

func factorialRecursive(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func main() {
	n := 4

	fmt.Printf("Факториал числа %d (итеративный) %d\n", n, factorialIterative(n))
	fmt.Printf("Факториал числа %d (рекурсивный) %d\n", n, factorialRecursive(n))
}
