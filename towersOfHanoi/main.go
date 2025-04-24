package main

import "fmt"

func hanoi(n int, A, B, C string) {
	if n == 1 {
		fmt.Printf("Move %d %s to %s \n", n, A, C)
	} else {
		hanoi(n-1, A, C, B) // Move the n-1th disc on the A through C to B
		fmt.Printf("Move %d from %s to %s \n", n, A, C)
		hanoi(n-1, B, A, C) // Move the n-1th disc on the B through A to C
	}
}

func main() {
	fmt.Print("Please enter the number of discs: \n")
	var n int
	fmt.Scanf("%d", &n)
	hanoi(n, "A", "B", "C")
}
