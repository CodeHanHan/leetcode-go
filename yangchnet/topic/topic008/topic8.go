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

type Automation struct {
	sign  int
	ans   int
	state string
	table map[string][]string
}

func (a *Automation) get(c byte) {
	a.state = a.table[a.state][a.get_col(c)]
	if a.state == "numed" {
		a.ans = a.ans*10 + int(c-'0')
		if a.sign == 1 {
			a.ans = min(a.ans, math.MaxInt32)
		} else {
			a.ans = min(a.ans, -math.MinInt32)
		}
	} else if a.state == "signed" {
		if c == '+' {
			a.sign = 1
		} else {
			a.sign = -1
		}
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func (a *Automation) get_col(c byte) int {
	switch c {
	case ' ':
		return 0
	case '+', '-':
		return 1
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return 2
	default:
		return 3
	}
}

func myAtoi1(s string) int {
	a := &Automation{
		sign:  1,
		ans:   0,
		state: "start",
		table: map[string][]string{
			"start":  {"start", "signed", "numed", "end"},
			"signed": {"end", "end", "numed", "end"},
			"numed":  {"end", "end", "numed", "end"},
			"end":    {"end", "end", "end", "end"},
		},
	}

	for i := 0; i < len(s); i++ {
		a.get(s[i])
	}

	return int(a.sign * a.ans)
}
