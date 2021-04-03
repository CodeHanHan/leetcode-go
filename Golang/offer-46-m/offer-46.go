package offer_46

import "strconv"

func translateNum(num int) int {
	numStr := strconv.Itoa(num)
	count := 0
	var f func(string, int)
	f = func(numStr string, i int) {
		if i >= len(numStr) {
			count += 1
			return
		}
		f(numStr, i+1)
		if i+1 < len(numStr) {
			if numStr[i] == '0' {
				return
			}
			if cur, _ := strconv.Atoi(numStr[i : i+2]); cur < 26 {
				f(numStr, i+2)
			}
		}
	}
	f(numStr, 0)
	return count
}

func translateNumDP(num int) int {
	numStr := strconv.Itoa(num)
	p, q, r := 0, 0, 1 //dp[i-2], dp[i-1], dp[i]
	for i := 0; i < len(numStr); i++ {
		p, q, r = q, r, 0 // 更新dp
		r += q            // dp[i] += dp[i-1]
		if i == 0 {
			continue
		}
		pre := numStr[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			r += p // dp[i] += dp[i-2]
		}
	}
	return r
}
