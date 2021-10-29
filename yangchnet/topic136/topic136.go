package topic136

func singleNumber(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] = nums[0] ^ nums[i]
	}
	return nums[0]
}
