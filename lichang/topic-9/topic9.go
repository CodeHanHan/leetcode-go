package topic9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	nums := make([]int, 0, 11)
	for x != 0 {
		nums = append(nums, x%10)
		x = x / 10
	}

	l := len(nums)
	for i := 0; i < l/2; i++ {
		if nums[i] != nums[l-i-1] {
			return false
		}
	}
	return true
}

func isPalindrome_1(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	y := 0
	for x > y {
		y = y*10 + x%10
		x /= 10
	}

	return x == y || x == y/10
}
