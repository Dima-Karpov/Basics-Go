package main

import "fmt"

const MAX_SIZE = 8
const MIN_SIZE = 1

type Fruit struct {
	name  string
	size  int
	price int
}

func main() {
	var item = [MAX_SIZE + 1]int{0}
	var value = [MAX_SIZE + 1]int{0}
	var fruits = [5]Fruit{
		{"Plum", 4, 4500},
		{"Apple", 5, 5700},
		{"Orange", 2, 2250},
		{"Strawberry", 1, 1100},
		{"Melon", 6, 6700},
	}

	var length = len(fruits)
	for i := 0; i < length; i++ {
		for j := fruits[i].size; j <= MAX_SIZE; j++ {
			var p = j - fruits[i].size
			var newValue = value[p] + fruits[i].price
			// Find the best solution
			if newValue > value[i] {
				value[j] = newValue
				item[j] = i
			}
		}
	}

	fmt.Printf("Item \t Price \t Size \n")
	for i := MAX_SIZE; i >= MIN_SIZE; i = i - fruits[item[i]].size {
		fmt.Printf("%s \t %d \t %d \n", fruits[item[i]].name, fruits[item[i]].price, fruits[item[i]].size)
	}

	fmt.Printf("\n")
	fmt.Printf("%d", value)
	fmt.Printf("Total \t %d \n", value[MAX_SIZE])
}
