package lc389

func findTheDifference1(s string, t string) byte {
	num := [26]int{}
	for _, ch := range s {
		num[ch-'a']++
	}
	var res byte
	for i, ch := range t {
		res = t[i]
		num[ch-'a']--
		if num[ch-'a'] < 0 {
			break
		}
	}
	return res
}

func findTheDifference2(s string, t string) byte {
	res := byte(0)
	for i := range s {
		res ^= s[i]
	}
	for i := range t {
		res ^= t[i]
	}
	return res
}
