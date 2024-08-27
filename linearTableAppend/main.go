package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)

	// create a new array
	var tempArr = make([]int, length+1)

	for i := 0; i < length; i++ {
		tempArr[i] = scores[i]
	}
	tempArr[length] = 75
	scores = tempArr

	for i := 0; i < length+1; i++ {
		fmt.Printf("%d\t", scores[i])
	}
}
