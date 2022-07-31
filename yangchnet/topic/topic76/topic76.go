package topic76

import "math"

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	len := math.MaxInt32
	ansL, ansR := -1, -1

	check := func() bool {
		for k, v := range ori {
			if cnt[k] < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < sLen; r++ {
		if r < sLen && ori[s[r]] > 0 {
			cnt[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < len {
				len = r - l + 1
				ansL, ansR = l, l+len
			}
			if _, ok := ori[s[l]]; ok {
				cnt[s[l]] -= 1
			}
			l++
		}
	}
	if ansL == -1 {
		return ""
	}
	return s[ansL:ansR]
}

func minWindow1(s string, t string) string {
	// 存储s和t中字符的出现次数
	hs, ht := make(map[byte]int), make(map[byte]int)

	// 记录t中字符出现次数
	for i, _ := range t {
		ht[t[i]]++
	}

	var res string
	for i, j, cnt := 0, 0, 0; i < len(s); i++ {
		hs[s[i]]++
		if hs[s[i]] <= ht[s[i]] { // 记录已经出现的有效字符
			cnt++
		}
		for j < len(s) && hs[s[j]] > ht[s[j]] { // s[j]变成多余的了
			hs[s[j]]--
			j++
		}
		if cnt == len(t) && (res == "" || i-j+1 < len(res)) {
			res = s[j : i+1]
		}
	}

	return res
}
