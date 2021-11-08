package topic003

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 0 {
		return 0
	}

	present := make(map[byte]int)
	present[s[0]] = 0

	start, maxLen := 0, 1
	max := 1

	for i := 1; i < len(s); i++ {
		index, ok := present[s[i]]
		if !ok || index < start {
			maxLen++
			if max < maxLen {
				max = maxLen
			}
		} else {
			start = index + 1
			maxLen = i - start + 1
		}
		present[s[i]] = i
	}

	return max
}
