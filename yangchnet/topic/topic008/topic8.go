package topic8

import (
	"math"
)

// 恶心解法
func myAtoi(s string) int {
	nagative := false
	var ptr int = 0
	for i, v := range s {
		if v == ' ' {
			continue
		} else {
			ptr = i
			break
		}
	}

	for i := ptr; i < len(s); i++ {
		if s[i] == '-' {
			ptr = i + 1
			nagative = true
			break
		} else if s[i] == '+' {
			ptr = i + 1
			nagative = false
			break
		} else {
			break
		}
	}

	var buf []int = make([]int, 0)
	for i := ptr; i < len(s); i++ {
		if int(s[i]) >= 48 && int(s[i]) <= 57 {
			buf = append(buf, int(s[i])-48)
			continue
		} else {
			break
		}
	}
	var ret int = 0

	var ok bool
	for i := 0; i < len(buf); i++ {
		ret, ok = Add32(ret*10, int(buf[i]))
		if !ok {
			if nagative {
				return -1 << 31
			} else {
				return 1<<31 - 1
			}
		}
	}

	if nagative {
		ret = -ret
	}

	if ret < -(1 << 31) {
		ret = -(1 << 31)
	}

	if ret > 1<<31-1 {
		ret = 1<<31 - 1
	}

	return ret
}

func Add32(left, right int) (int, bool) {
	if right > 0 {
		if left > math.MaxInt32-right {
			return 0, false
		}
	} else {
		if left < math.MinInt32-right {
			return 0, false
		}
	}
	return left + right, true
}
