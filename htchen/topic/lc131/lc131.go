package lc131

func partition(s string) [][]string {
	n := len(s)
	var ans [][]string
	var path []string
	var trackback func(start int)
	trackback = func(start int) {
		if start == n {
			temp := make([]string, len(path))
			copy(temp, path)
			ans = append(ans, temp)
			return
		}
		for i := start; i < n; i++ {
			if isPalindrome(s[start : i+1]) {
				path = append(path, s[start:i+1])
			} else {
				continue
			}
			trackback(i + 1)
			path = path[:len(path)-1]
		}
	}
	trackback(0)
	return ans
}

func isPalindrome(s string) bool {
	n := len(s)
	for i, j := 0, n-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
