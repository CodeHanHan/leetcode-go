package offer_63_m

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	dp := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if dp > prices[i]-minPrice {
			dp = dp
		} else {
			dp = prices[i] - minPrice
		}
		if minPrice > prices[i] {
			minPrice = prices[i]
		}
	}
	return dp
}
