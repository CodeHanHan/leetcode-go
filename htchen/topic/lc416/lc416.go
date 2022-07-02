package lc416

func canPartition(nums []int) bool {
	n := len(nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}

	target := sum / 2
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, target+1)
	}
	if nums[0] <= target {
		dp[0][nums[0]] = true
	}
	for i := 1; i < n; i++ {
		for j := 0; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if nums[i] == j { 
				dp[i][j] = true
				continue
			}
			if nums[i] < j {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			}
		}
	}
	return dp[n-1][target]
}
