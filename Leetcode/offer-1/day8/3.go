package day8

import (
	"math"
)

func maxProfit(prices []int) int {
	minPrice := math.MaxInt64
	maxPrice := 0

	for i := 0; i < len(prices); i++ {
		maxPrice = max(prices[i]-minPrice, maxPrice)
		minPrice = min(prices[i], minPrice)
	}

	return maxPrice
}

func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func MaxProfit(prices []int) int {
	return maxProfit(prices)
}
