package goden1709

import "math"

func getKthMagicNumber(k int) int {
	array := make([]int, k)
	array[0] = 1
	var p3, p5, p7 int = 0, 0, 0
	for i := 1; i < k; i++ {
		p := minN(array[p3]*3, array[p5]*5, array[p7]*7)
		if p%3 == 0 {
			p3++
		}
		if p%5 == 0 {
			p5++
		}
		if p%7 == 0 {
			p7++
		}
		array[i] = p
	}

	return array[k-1]
}

func minN(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if min > num {
			min = num
		}
	}

	return min
}
