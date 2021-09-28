package topic5

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([]int, len(s))
	dp[0] = 1

	var start int = 0 // 指向最长回文串的开始元素

	for i := 1; i < len(s); i++ {
		l := -1
		for j := 0; j < i; j++ {
			if s[j] == s[i] && isPalindrome(s[j:i+1]) {
				l = max(l, i-j+1)
			}
		}

		if dp[i-1] >= l {
			dp[i] = dp[i-1]
		} else {
			dp[i] = l
			start = i - l + 1
		}
	}

	return s[start : start+dp[len(dp)-1]]
}

func isPalindrome(s string) bool {
	var i, j int = 0, len(s) - 1
	for i <= j && s[i] == s[j] {
		i++
		j--
	}
	return i > j
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// ==========================================

func longestPalindrome_1(s string) string {
	if len(s) <= 1 {
		return s
	}

	cur := 1 // dp[i-1], dp[i] , 初始化dp[0] = 1

	var start int = 0 // 指向最长回文串的开始元素

	for i := 1; i < len(s); i++ {
		pre := cur

		l := -1 // 以下标i为结尾的字符串的最长回文子串
		for j := 0; j < i; j++ {
			if s[j] == s[i] && isPalindrome(s[j:i+1]) {
				l = max(l, i-j+1)
			}
		}

		if pre >= l {
			cur = pre
		} else {
			cur = l
			start = i - l + 1
		}
	}

	return s[start : start+cur]
}
