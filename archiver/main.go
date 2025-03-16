package main

import (
	"fmt"
	"strconv"
)

func archiveString(input string) string {
	archived := ""
	count := 1

	for i := 1; i < len(input); i++ {
		if input[i] == input[i-1] {
			count++
		}
		archived += string(input[i-1]) + strconv.Itoa(count)
		count = 1
	}
	archived += string(input[len(input)-1]) + strconv.Itoa(count)

	return archived
}

func main() {
	str := "abccccc"

	archived := archiveString(str)
	fmt.Printf("archived %v\n", archived)

}
