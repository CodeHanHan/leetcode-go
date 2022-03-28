package topic121

import "math"

// 一次遍历
func maxProfit(prices []int) int {
	var minPrice int = math.MaxInt32
	var profit int = 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > profit {
			profit = prices[i] - minPrice
		}
	}

	return profit
}
