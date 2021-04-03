package offer_48_m_

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	set := make(map[uint8]int)
	dp, max := 1, 1 // p: dp[i]
	set[s[0]] = 1   // 位置从1开始，位置为0表示不存在
	for i := 1; i < len(s); i++ {
		if set[s[i]] <= 0 { // 当前字符在子串中不存在
			dp += 1
			set[s[i]] = i + 1
		} else { // 当前字符在子串中存在
			j := i - dp
			for s[j] != s[i] {
				set[s[j]] = 0
				j++
			}
			dp = i - j
			set[s[i]] = i + 1
		}
		if max < dp {
			max = dp
		}
	}
	return max
}
