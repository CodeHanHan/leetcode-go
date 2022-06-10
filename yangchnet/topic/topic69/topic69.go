package topic69

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}

	l, r := 1, x/2
	for l <= r {
		mid := l + (r-l)/2
		if mid*mid == x {
			return mid
		} else if mid*mid < x {
			l = mid + 1
		} else if mid*mid > x {
			r = mid - 1
		}
	}

	return r
}
