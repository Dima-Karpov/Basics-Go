package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)
	var tempArr = make([]int, length+1)

	insert(scores, length, tempArr, 75, 2)
	scores = tempArr

	for i := 0; i < length+1; i++ {
		fmt.Printf("%d\t", scores[i])
	}
}

func insert(array []int, length int, tempArr []int, score int, insertIndex int) {
	for i := 0; i < length; i++ {
		if i < insertIndex {
			tempArr[i] = array[i]
		} else {
			tempArr[i+1] = array[i]
		}
	}

	tempArr[insertIndex] = score
}
