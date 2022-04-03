package goden1611

// 递归解法，时间超时
func divingBoard(shorter int, longer int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	res := make([]int, 0)
	existd := make(map[int]bool)

	var fn func(k int, length int)
	fn = func(k int, length int) {
		if k <= 0 {
			if _, ok := existd[length]; !ok {
				res = append(res, length)
				existd[length] = true
			}
			return
		}

		fn(k-1, length+shorter)
		fn(k-1, length+longer)
	}

	fn(k, 0)

	// sort.Ints(res)
	return res
}

func divingBoard2(shorter int, longer int, k int) []int {
	if k <= 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{longer * k}
	}

	res := make([]int, 0)
	for i := k * shorter; i <= k*longer; i += (longer - shorter) {
		res = append(res, i)
	}

	return res
}
