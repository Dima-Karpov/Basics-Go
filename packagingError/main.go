package main

import "fmt"

func main() {
	i, err := someFunc(0)
	fmt.Println(i, err)

	i, err = someFunc(10)
	fmt.Println(i, err)
}

func someFunc(i int) (int, error) {
	j, err := funcReturningError(i)
	if err != nil {
		return 0, fmt.Errorf("wrap error: %w", err)
	}
	return j, nil
}

func funcReturningError(i int) (int, error) {
	if i == 0 {
		return 0, fmt.Errorf("got %d", i)
	}
	return i, nil
}
