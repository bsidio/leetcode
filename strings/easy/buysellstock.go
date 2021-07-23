package main

import "fmt"

func main() {
	prices := []int{2, 4, 1, 3, 5, 7}

	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	fp := prices[0]
	mp := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < fp {
			fp = prices[i]

		} else {
			if prices[i]-fp > mp {
				mp = prices[i] - fp
			}
		}
	}

	return mp
}
