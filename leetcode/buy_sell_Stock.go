package main

import (
	"fmt"
)

func maxProfit(price []int) int {
	// max_profit_ind := make(map[int]int)
	// var maxProfit, lowest int
	lowest := price[0]
	maxProfit := 0
	for _, value := range price {
		if value < lowest {
			lowest = value
		}
		profit := value - lowest
		if maxProfit < profit {
			maxProfit = profit
		}
	}

	return maxProfit
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}

	fmt.Println(maxProfit(prices))
}
