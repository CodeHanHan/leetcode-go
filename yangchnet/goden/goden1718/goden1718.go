package goden1718

func shortestSeq(big []int, small []int) []int {
	n := len(big)
	need := make(map[int]int) // need存储small中各元素的个数
	minLen, diff := n, 0

	for _, v := range small {
		need[v]++
		diff++
	}

	var res []int = []int{}
	l, r := 0, 0
	for ; r < n; r++ {
		if _, ok := need[big[r]]; ok {
			need[big[r]]--
			if need[big[r]] >= 0 { // need[big[r]] >= 0, 说明big[r]是有效的字符
				diff--
			}
		}

		for diff == 0 {
			if r-l < minLen {
				minLen = r - l
				res = []int{l, r}
			}
			if _, ok := need[big[l]]; ok {
				need[big[l]]++
				if need[big[l]] > 0 {
					diff++
				}
			}
			l++
		}
	}

	return res

}
