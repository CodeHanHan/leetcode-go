package goden1711

import "math"

func findClosest(words []string, word1 string, word2 string) int {
	var p1, p2 int = -1, -1

	min := math.MaxInt64
	for i, word := range words {
		switch word {
		case word1:
			p1 = i
		case word2:
			p2 = i
		default:
			continue
		}
		if p1 > 0 && p2 > 0 {
			min = minN(min, abs(p1-p2))
		}
	}

	return min
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

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}
