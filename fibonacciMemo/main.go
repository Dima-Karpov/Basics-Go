package main

import "fmt"

func fibonacciMemo(n int) int {
	memo := [2]int{1, 1} // f(1) = 1, f(2) = 1
	for i := 2; i <= n; i++ {
		memo[i%2] = memo[0] + memo[1]
	}

	return memo[(n+1)%2]
}

func main() {
	fmt.Println("F(10):", fibonacciMemo(10)) // 55
	fmt.Println("F(5):", fibonacciMemo(5))   // 5
	fmt.Println("F(1):", fibonacciMemo(1))   // 1
}
