package topic209

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := nums[0]

	var min int = len(nums) + 1
	for right < len(nums) {
		if nums[right] >= target {
			return 1
		}
		if sum >= target {
			if right-left+1 < min {
				min = right - left + 1
			}
			sum -= nums[left]
			left++
		} else if sum < target {
			right++
			if right < len(nums) {
				sum += nums[right]
			}
		}
	}

	if min > len(nums) {
		return 0
	}

	return min
}
