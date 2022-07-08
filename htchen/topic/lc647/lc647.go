package lc647

func countSubstrings(s string) int {
	n := len(s)
	num := 0
	for i := 0; i < n; i++ {
		for j := 0; j <= 1; j++ {
			l, r := i, i+j
			for l >= 0 && r < n && s[l] == s[r] {
				num++
				l--
				r++
			}
		}
	}
	return num
}
