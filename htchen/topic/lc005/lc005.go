package lc005

func longestPalindrome1(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}

	maxLen := 1
	begin := 0
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for L := 2; L <= n; L++ {
		for i := 0; i < n; i++ {
			j := L + i - 1
			if j >= n {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && L > maxLen {
				maxLen = L
				begin = i
			}
		}
	}
	return s[begin : begin+maxLen]
}

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
