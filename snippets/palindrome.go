package snippets

// 回文串相关

// 判断是否是回文串
func isPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			break
		}

		l++
		r--
	}

	return l >= r
}

// 中心扩展寻找回文串, 返回回文串的左右边界
func expendCenter(s string, i, j int) (int, int) {
	for ; i >= 0 && j < len(s); i, j = i-1, j+1 {
		if s[i] != s[j] {
			break
		}
	}

	return i + 1, j - 1
}
