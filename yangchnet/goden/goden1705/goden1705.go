package goden1705

func findLongestSubarray(array []string) []string {
	m := map[int]int{} // 字母和数字的个数差是某个数的最小位置

	cntAlpha, cntNum := 0, 0
	start, end := -1, -1
	maxLen := 0
	m[0] = -1
	for idx, s := range array {

		if s[0] >= '0' && s[0] <= '9' {
			cntNum++
		} else {
			cntAlpha++
		}

		diff := cntAlpha - cntNum
		if pos, ok := m[diff]; ok {
			if idx-pos > maxLen {
				start = pos
				end = idx
				maxLen = idx - pos
			}
		} else {
			m[diff] = idx
		}
	}

	return array[start+1 : end+1]
}
