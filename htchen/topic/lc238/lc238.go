package lc238

func productExceptSelf1(nums []int) []int {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	ans := make([]int, n)

	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	right[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		ans[i] = left[i] * right[i]
	}
	return ans
}

func productExceptSelf2(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)

	ans[0] = 1
	for i := 1; i < n; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] = ans[i] * right
		right *= nums[i]
	}
	return ans
}
