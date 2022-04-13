package goden1710

func majorityElement(nums []int) int {
	curNum := nums[0]
	tickets := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == curNum {
			tickets++
		} else {
			tickets--
			if tickets < 0 {
				curNum = nums[i]
				tickets = 1
			}
		}
	}

	if tickets == 0 {
		return -1
	}

	tickets = 0
	for _, num := range nums {
		if num == curNum {
			tickets++
		}
	}

	if tickets <= len(nums)/2 {
		return -1
	}

	return curNum
}
