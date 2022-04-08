package goden1704

// 使用哈希表
func missingNumber(nums []int) int {
	exist := make(map[int]bool)

	for i := 0; i <= len(nums); i++ {
		exist[i] = true
	}

	for _, n := range nums {
		delete(exist, n)
	}

	for k, _ := range exist {
		return k
	}

	return -1
}

// 求和再相减
func missingNumber2(nums []int) int {
	n := len(nums)
	sum := n * (n + 1) / 2
	for _, num := range nums {
		sum -= num
	}

	return sum
}

// 位运算
func missingNumber3(nums []int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		res ^= i
		res ^= nums[i]
	}

	res ^= len(nums)

	return res
}
