package topic300

import "sort"

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	dp := make([]int, len(nums))
	dp[0] = 1

	var res int = dp[0]
	for i := 1; i < len(nums); i++ {
		m := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				m = max(m, dp[j])
			}
		}
		dp[i] = m + 1
		if res < dp[i] {
			res = dp[i]
		}
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLIS2(nums []int) int {
	dp := make([]int, len(nums))

	maxL := 0
	for _, num := range nums {
		idx := sort.Search(maxL, func(i int) bool {
			return dp[i] >= num
		})
		dp[idx] = num
		if idx == maxL {
			maxL++
		}
	}

	return maxL
}
