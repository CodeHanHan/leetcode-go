package goden1708

import "sort"

func bestSeqAtIndex(height []int, weight []int) int {
	slice := make([][]int, len(height))
	for i, _ := range slice {
		slice[i] = []int{height[i], weight[i]}
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i][0] < slice[j][0] { // 先按身高从小到大排
			return true
		}
		if slice[i][0] > slice[j][0] {
			return false
		}
		return slice[i][1] > slice[j][1] // 再按体重倒序从大到小，防止同样身高的人重复放置
	})

	dp := make([]int, len(height))
	dp[0] = 1

	res := dp[0]
	for i := 1; i < len(height); i++ {
		m := 0
		for j := 0; j < i; j++ {
			if slice[j][1] < slice[i][1] {
				m = max(m, dp[j])
			}
		}

		dp[i] = m + 1
		res = max(res, dp[i])
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}

func bestSeqAtIndex2(height []int, weight []int) int {
	slice := make([][]int, len(height))
	for i, _ := range slice {
		slice[i] = []int{height[i], weight[i]}
	}

	sort.Slice(slice, func(i, j int) bool {
		if slice[i][0] < slice[j][0] { // 先按身高从小到大排
			return true
		}
		if slice[i][0] > slice[j][0] {
			return false
		}
		return slice[i][1] > slice[j][1] // 再按体重倒序从大到小，防止同样身高的人重复放置
	})

	dp := make([]int, len(height))

	maxL := 0
	for _, s := range slice {
		idx := sort.Search(maxL, func(i int) bool {
			return dp[i] >= s[1]
		})
		dp[idx] = s[1]
		if idx == maxL {
			maxL++
		}
	}

	return maxL
}
