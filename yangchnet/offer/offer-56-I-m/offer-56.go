package offer_56_I_m

func singleNumbers(nums []int) []int {
	ret := 0
	for _, num := range nums {
		ret ^= num
	}

	mask := 1
	for ret & mask == 0 {
		mask = mask << 1
	}

	a, b := 0, 0
	for _, num := range nums {
		if num & mask == 0 {
			a ^= num
		} else {
			b ^= num
		}
	}
	return []int{a, b}
}
