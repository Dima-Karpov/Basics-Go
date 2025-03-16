package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(nameShuffter("john McClane"))

}

func nameShuffter(str string) string {
	names := strings.Fields(str)
	if len(names) > 1 {
		return names[1] + " " + names[0]
	}

	return str

}
