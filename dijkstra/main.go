package main

import "fmt"

func main() {
	var flags = []string{"R", "B", "W", "B", "R", "W"}
	var length = len(flags)
	var b, w, r, count = 0, 0, length - 1, 0
	for {
		if w > r {
			break
		}
		if flags[w] == "W" {
			w++
		} else if flags[w] == "B" {
			var temp = flags[w]
			flags[b] = temp
			w++
			b++
			count++
		} else if flags[w] == "R" {
			var m = flags[w]
			flags[w] = flags[r]
			flags[r] = m
			r--
			count++
		}
	}
	for i := 0; i < length; i++ {
		fmt.Printf("%s", flags[i])
	}

	fmt.Printf("\nThe total exchange count: %d", count)

}
