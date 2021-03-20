package main

import "fmt"

func maxProfit(prices []int) int {
	maxProfit := 0
	maxPriceForEachDateMap := make(map[int]int, len(prices))
	maxPrice := 0

	for i := len(prices) - 1; i >= 0; i-- {
		if maxPrice < prices[i] {
			maxPrice = prices[i]
		}

		maxPriceForEachDateMap[i] = maxPrice
	}

	for i := 0; i < len(prices); i++ {
		maxProfitIfBuy := maxPriceForEachDateMap[i] - prices[i]
		if maxProfitIfBuy > maxProfit {
			maxProfit = maxProfitIfBuy
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
