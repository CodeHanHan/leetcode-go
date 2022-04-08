package goden1621

func findSwapValues(array1 []int, array2 []int) []int {
	lSum, rSum := 0, 0
	m2 := make(map[int]bool)
	for _, v := range array1 {
		lSum += v
	}

	for _, v := range array2 {
		rSum += v
		m2[v] = true
	}

	if (lSum-rSum)%2 != 0 { // 差值不为偶
		return []int{}
	}

	for _, l := range array1 {
		if m2[l-(lSum-rSum)/2] {
			return []int{l, l - (lSum-rSum)/2}
		}
	}

	return []int{}
}
