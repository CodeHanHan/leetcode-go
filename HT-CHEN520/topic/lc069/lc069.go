package lc069

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	left, right := 1, x
	res := 0
	for left <= right {
		mid := (left + right) / 2
		if x/mid > mid {
			left = mid + 1
			res = mid
		} else if x/mid == mid {
			return mid
		} else {
			right = mid - 1
		}
	}
	return res
}
