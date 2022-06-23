package topic75

func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}

	left, right := 0, len(nums)-1

	for left < len(nums) && nums[left] == 0 {
		left++
	}

	for right >= 0 && nums[right] == 2 {
		right--
	}

	for i := left; i <= right; {
		switch nums[i] {
		case 0:
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		case 1:
			i++
			continue
		case 2:
			nums[right], nums[i] = nums[i], nums[right]
			right--
		}
	}

	return
}
