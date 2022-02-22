package topic238

func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	r := 1
	for i := 0; i < len(nums); i++ {
		left[i] = r
		r *= nums[i]
	}

	r = 1
	for i := len(nums) - 1; i >= 0; i-- {
		right[i] = r
		r *= nums[i]
	}

	for i := 0; i < len(nums); i++ {
		left[i] = left[i] * right[i]
	}

	return left
}

// 优化
func productExceptSelf_1(nums []int) []int {
	tmp := make([]int, len(nums))

	r := 1
	for i := 0; i < len(nums); i++ {
		tmp[i] = r
		r *= nums[i]
	}

	r = 1
	var k int
	for i := len(nums) - 1; i >= 0; i-- {
		k = nums[i]
		nums[i] = r
		r *= k

		tmp[i] = tmp[i] * nums[i]
	}

	return tmp
}
