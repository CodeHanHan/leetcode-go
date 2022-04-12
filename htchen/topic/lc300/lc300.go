package lc300

func lengthOfLIS1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}
	return max(dp...)
}

func max(a ...int) int {
	res := a[0]
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}

func lengthOfLIS2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	d := make([]int, n+1)
	len := 1
	d[len] = nums[0]
	for i := 0; i < n; i++ {
		if nums[i] > d[len] {
			len += 1
			d[len] = nums[i]
		} else {
			left, right, pos := 1, len, 0
			for left <= right {
				mid := (left + right) / 2
				if nums[i] > d[mid] {
					pos = mid
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			d[pos+1] = nums[i]
		}
	}
	return len
}
