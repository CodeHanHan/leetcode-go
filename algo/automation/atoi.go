package automation

import "math"

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

func myAtoi(s string) int {
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
