package topic5

// 动态规划原始版本
func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([]int, len(s))
	dp[0] = 1

	var start int = 0 // 指向最长回文串的开始元素

	for i := 1; i < len(s); i++ {
		l := -1                  // 回文串长度
		for j := 0; j < i; j++ { // 找寻以当前坐标为结束的字符串的最长回文串长度
			if s[j] == s[i] && isPalindrome(s[j:i+1]) {
				l = max(l, i-j+1)
			}
		}

		if dp[i-1] >= l { // 翻译状态转移方程
			dp[i] = dp[i-1]
		} else {
			dp[i] = l
			start = i - l + 1
		}
	}

	return s[start : start+dp[len(dp)-1]]
}

//是否是回文串
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

// 优化了dp数组的动态规划版本
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

// ==========================================

// 中心扩展法
func longestPalindrome_2(s string) string {
	var start, l int = 0, 1 // start 存储回文串的开始位置下标，l存储最长回文串的长度

	for k := 0; k < len(s); k++ {
		var i, j int
		for i = k; i >= 0; i-- { // 向左扩展
			if s[i] != s[k] {
				break
			}
		}

		for j = k; j < len(s); j++ { // 向右扩展
			if s[j] != s[k] {
				break
			}
		}

		for i >= 0 && j < len(s) && s[i] == s[j] { // 两边扩展
			i--
			j++
		}

		// 恢复i和j的值
		i++
		j--

		if j-i+1 > l { // 更新start和l
			l = j - i + 1
			start = i
		}
	}
	return s[start : start+l]
}

// ==========================================

// 中心扩展法2
func longestPalindrome2(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expendCenter(s, i, i)
		left2, right2 := expendCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expendCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s); left, right = left-1, right+1 {
		if s[left] != s[right] {
			break
		}
	}
	return left + 1, right - 1
}
