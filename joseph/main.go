package main

import "fmt"

const N = 11
const M = 3

func main() {
	var man = [N]int{0}
	fmt.Printf("man = %v\n", man)

	var count = 1
	var i = 0
	var pos = -1

	for {
		if count > N {
			break
		}
		for {
			pos = (pos + 1) % N // Ring

			fmt.Printf("pos = %v\n", pos)

			if man[pos] == 0 {
				i++
			}
			if i == M {
				i = 0
				break
			}
		}
		man[pos] = count
		count++
	}
	fmt.Printf("\nJoseph sequence : ")
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", man[i])
	}
}
