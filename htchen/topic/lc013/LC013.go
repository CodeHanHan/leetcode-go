package LC013

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	n := len(s)
	ans := m[s[n-1]]
	for i := n - 2; i >= 0; i-- {
		if m[s[i]] >= m[s[i+1]] {
			ans += m[s[i]]
		} else {
			ans -= m[s[i]]
		}
	}
	return ans
}
