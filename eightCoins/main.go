package main

import (
	"fmt"
	"math/rand"
)

func compare(coins []int, i, j, k int) { // coin[k] true, coin[i] > coin[j]
	// coin[i]>coin[j]&&coin[i]>coin[k]---->coin[i]
	// is a heavy counterfeit coin
	if coins[i] > coins[k] {
		fmt.Printf("\nCounterfeit currency %d is heavier\n", i+1)
	} else { // coin[j] is a light counterfeit coin
		fmt.Printf("\nCounterfeit currency %d is lighter\n", j+1)
	}
}

func detectCounterfeitCoin(coins []int) {
	// Implementing binary search approach to reduce comparisons
	leftSum := coins[0] + coins[1] + coins[2]
	rightSum := coins[3] + coins[4] + coins[5]

	switch {
	case leftSum == rightSum:
		if coins[6] > coins[7] {
			compare(coins, 6, 7, 0)
		} else {
			compare(coins, 7, 6, 0)
		}
	case leftSum > rightSum:
		if coins[0]+coins[3] == coins[1]+coins[4] { //(a+e)==(d+b)
			compare(coins, 2, 5, 0)
		} else if coins[0]+coins[3] > coins[1]+coins[4] { //(a+e)>(d+b)
			compare(coins, 0, 4, 1)
		} else {
			compare(coins, 1, 3, 0)
		}
	case leftSum < rightSum:
		if coins[0]+coins[3] == coins[1]+coins[4] { //(a+e)==(d+b)
			compare(coins, 5, 2, 0)
		} else if coins[0]+coins[3] > coins[1]+coins[4] { //(a+e)>(d+b)
			compare(coins, 3, 1, 0)
		} else {
			compare(coins, 4, 0, 1)
		}
	}
}

func main() {
	var coins = make([]int, 8)
	// Initial coin weight is 10
	for i := range coins {
		coins[i] = 10
	}

	// Randomly generate weight for counterfeit coin
	coin := rand.Intn(20) // Assuming max weight can be 20
	coinIndex := rand.Intn(8)
	coins[coinIndex] = coin

	detectCounterfeitCoin(coins)

	// Display final weights of coins
	for i := range coins {
		fmt.Printf("%d, ", coins[i])
	}
}
