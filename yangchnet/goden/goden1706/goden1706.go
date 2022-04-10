package goden1706

import "strconv"

// 暴力法
func numberOf2sInRange(n int) int {
	if n < 2 {
		return 0
	}
	cnt := 0
	for i := 2; i <= n; i++ {
		s := strconv.Itoa(i)
		for j := 0; j < len(s); j++ {
			if s[j] == '2' {
				cnt++
			}
		}
	}

	return cnt
}

// 统计法
func numberOf2sInRange2(n int) int {
	cnt := 0

	s := strconv.Itoa(n)
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i] - '0'
		if c > 2 {
			l, _ := strconv.Atoi(s[0:i])
			cnt += (l + 1) * pow(10, len(s)-i-1)
		} else if c < 2 {
			l, _ := strconv.Atoi(s[0:i])
			cnt += l * pow(10, len(s)-i-1)
		} else if c == 2 {
			l, _ := strconv.Atoi(s[0:i])
			r, _ := strconv.Atoi(s[i+1:])
			cnt += l * pow(10, len(s)-i-1)
			cnt += r + 1
		}
	}

	return cnt
}

// pow return x**y
func pow(x, y int) int {
	ret := 1
	for i := 0; i < y; i++ {
		ret *= x
	}

	return ret
}
