package topic231

import "math"

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}

	var l, r int = 0, n
	if r > 31 {
		r = 31
	}

	var mid, pow int
	for l <= r {
		mid = l + (r-l)/2
		pow = int(math.Pow(float64(2), float64(mid)))
		if pow == n {
			return true
		} else if pow < n {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
