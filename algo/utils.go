package algo

// MaxN returns the maximum number of elements
func MaxN(vs ...int) int {
	max := vs[0]
	for _, v := range vs {
		if max < v {
			max = v
		}
	}

	return max
}

// MinN return minest number of elements
func MinN(vs ...int) int {
	min := vs[0]
	for _, v := range vs {
		if v < min {
			min = v
		}
	}

	return min
}
