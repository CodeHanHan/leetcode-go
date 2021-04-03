package offer_39_s_

func majorityElement(nums []int) int {
	cur := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if cur == nums[i] {
			count++
			continue
		} else {
			if count > 0 {
				count--
				continue
			}
			cur = nums[i]
			count = 1
		}
	}
	return cur
}
