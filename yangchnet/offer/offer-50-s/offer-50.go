package offer_50_s_

func firstUniqChar(s string) byte {
	dict := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		dict[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if dict[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
