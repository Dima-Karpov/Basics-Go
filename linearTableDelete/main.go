package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	fmt.Printf("Please enter the index to be deleted: \n")
	var index int
	fmt.Scan(&index)

	var length = len(scores)
	var tempArr = make([]int, length-1)

	for i := 0; i < length; i++ {
		if i < index {
			tempArr[i] = scores[i]
		}

		if i > index {
			tempArr[i-1] = scores[i]
		}
	}
	scores = tempArr

	for i := 0; i < length-1; i++ {
		fmt.Printf("%d\t", scores[i])
	}
}
