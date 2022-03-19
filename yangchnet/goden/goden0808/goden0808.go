package goden0808

import "sort"

func permutation(S string) []string {
	res := make([]string, 0)
	visit := make([]bool, len(S))
	a := []byte(S)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	S = string(a)
	var backTrack func(cur string)
	backTrack = func(cur string) {
		if len(cur) == len(S) {
			res = append(res, cur)
			return
		}
		for i := 0; i < len(S); i++ {
			if visit[i] || i > 0 && S[i-1] == S[i] && visit[i-1] && !visit[i] {
				continue
			}
			visit[i] = true
			backTrack(cur + string((S[i])))
			visit[i] = false
		}
	}
	backTrack("")
	return res
}
