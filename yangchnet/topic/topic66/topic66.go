package topic66

func plusOne(digits []int) []int {
	var c int = 1

	var v int
	for i := len(digits) - 1; i >= 0; i-- {
		v = digits[i]
		digits[i] = (digits[i] + c) % 10
		c = (v + c) / 10
	}

	if c != 0 {
		return append([]int{c}, digits...)
	}

	return digits
}
