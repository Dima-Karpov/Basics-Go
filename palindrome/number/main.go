package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	length := len(str)

	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	numbers := []int{101, 43134, 12345, 1221, 55555}

	for _, number := range numbers {
		if isPalindrome(number) {
			fmt.Printf("%d is a palindrome\n", number)
		} else {
			fmt.Printf("%d is not a palindrome\n", number)
		}
	}
}
