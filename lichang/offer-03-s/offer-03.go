package offer_03_s

func findRepeatNumber(nums []int) int {

	i := 0
	for {
		for nums[i] == i {
			i++
		}
		if nums[i] == nums[nums[i]] {
			return nums[i]
		}
		nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
	}
}
