package lc003

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	m := map[byte]int{}
	right, res := -1, 0
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}
		for right+1 < n && m[s[right+1]] == 0 {
			m[s[right+1]] = 1
			right++
		}
		res = max(res, right-i+1)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
