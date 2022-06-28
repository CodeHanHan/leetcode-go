package topic93

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ans := make([]string, 0)
	n := len(s)
	if n <= 3 || n >= 13 {
		return ans
	}

	vis := make([]bool, n)
	dots := []int{0}
	var bp func(idx int)
	bp = func(idx int) {
		if len(dots) == 4 {
			if ok, ip := valid(dots, s); ok {
				ans = append(ans, ip)
			}
			return
		}

		for i := idx; i < n; i++ {
			if vis[i] {
				continue
			}
			dots = append(dots, i)
			vis[i] = true
			bp(i + 1)
			vis[i] = false
			dots = dots[:len(dots)-1]
		}
	}

	bp(1)

	return ans
}

func valid(dots []int, s string) (ok bool, ip string) {
	dots = append(dots, len(s))
	numSlice := make([]int, 0)
	ipSlice := make([]string, 0)
	for i := 1; i <= 4; i++ {
		if dots[i]-dots[i-1] > 3 {
			return false, ""
		}

		num, err := strconv.Atoi(s[dots[i-1]:dots[i]])
		if err != nil || num < 0 || num > 255 {
			return false, ""
		}

		numSlice = append(numSlice, num)
		ipSlice = append(ipSlice, fmt.Sprintf("%d", num))
	}

	ip = strings.Join(ipSlice, ".")
	if len(ip)-3 != len(s) {
		return false, ""
	}
	return true, ip
}

func restoreIpAddresses1(s string) []string {
	ans := make([]string, 0)

	var bp func(s string, tmp []string)
	bp = func(s string, tmp []string) {
		if len(s) == 0 && len(tmp) == 4 {
			ans = append(ans, strings.Join(tmp, "."))
			return
		}

		if len(tmp) < 4 {
			l := min(3, len(s))
			for i := 0; i < l; i++ {
				p, n := s[:i+1], s[i+1:]
				pn, err := strconv.Atoi(p)
				if err != nil {
					return
				}
				if p != "" && pn >= 0 && pn <= 255 && fmt.Sprintf("%d", pn) == p {
					bp(n, append(tmp, p))
				}
			}
		}
	}

	bp(s, []string{})

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
