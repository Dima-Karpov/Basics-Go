package main

import (
	"bufio"
	"fmt"
	"os"
)

func Scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err := in.Err(); err != nil {
		fmt.Println("Some error occurred:", err)
	}

	return in.Text()
}

func Scan2() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Some error occurred:", err)
	}

	return str
}

func main() {
	//fmt.Println("Scan1: ", Scan1())
	fmt.Println("Scan2: ", Scan2())
}
