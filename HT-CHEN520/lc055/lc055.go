package lc055

func canJump(nums []int) bool {
	n := len(nums)
	reachMax := 0
	for i := 0; i < n; i++ {
		if i <= reachMax && i+nums[i] > reachMax {
			reachMax = i + nums[i]
		}
		if reachMax >= n-1 {
			return true
		}
	}
	return reachMax >= n-1
}
