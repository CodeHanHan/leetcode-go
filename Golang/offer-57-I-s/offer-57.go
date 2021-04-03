package offer_57_I_s

func twoSum(nums []int, target int) []int {
	var p, q int = 0, len(nums) - 1
	for p != q {
		sum := nums[p] + nums[q]
		if sum < target {
			p++
		} else if sum > target {
			q--
		} else {
			return []int{nums[p], nums[q]}
		}
	}
	return []int{}
}
