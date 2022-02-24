package lc121

func maxProfit(prices []int) int {
	minprice := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < minprice {
			minprice = prices[i]
		} else if prices[i]-minprice > maxProfit {
			maxProfit = prices[i] - minprice
		}
	}
	return maxProfit
}
