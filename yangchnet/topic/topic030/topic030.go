package topic030

import "reflect"

func all(words []string) map[string]bool {
	rets := make(map[string]bool)
	used := make([]bool, len(words))

	l := len(words) * len(words[0])
	var fn func(curStr string)
	fn = func(curStr string) {
		if len(curStr) == l {
			rets[curStr] = true
			return
		}

		for i := 0; i < len(words); i++ {
			if !used[i] {
				used[i] = true
				fn(curStr + words[i])
				used[i] = false
			}
		}
	}

	fn("")

	return rets
}

func findSubstring(s string, words []string) []int {
	n := len(words) * len(words[0])
	if len(s) < n {
		return []int{}
	}

	present := all(words)

	rets := make([]int, 0)
	for i := 0; i <= len(s)-n; i++ {
		if present[s[i:i+n]] {
			rets = append(rets, i)
		}
	}

	return rets
}

func findSubstring_1(s string, words []string) []int {
	oriMap := map[string]int{}
	for _, w := range words {
		oriMap[w]++
	}

	wordSize := len(words[0])
	stepSize := wordSize * len(words)
	var check = func(sk string) bool {
		nowMap := map[string]int{}
		for i := 0; i+wordSize <= len(sk); i += wordSize {
			wk := sk[i : i+wordSize]
			nowMap[wk]++
		}
		return reflect.DeepEqual(nowMap, oriMap)
	}

	res := []int{}
	for i := 0; i+stepSize <= len(s); i++ {
		ss := s[i : i+stepSize]
		if check(ss) {
			res = append(res, i)
		}
	}

	return res
}
