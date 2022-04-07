package lc128

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	longestlen := 0
	for num := range numSet {
		if !numSet[num-1] {
			currNum := num
			currLen := 1
			for numSet[currNum+1] {
				currNum++
				currLen++
			}
			if longestlen < currLen {
				longestlen = currLen
			}
		}
	}
	return longestlen
}
