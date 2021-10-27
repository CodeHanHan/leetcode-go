package LC014

func longestCommonPrefix(strs []string) string {
	strlen := len(strs)
	if strlen == 0 {
		return ""
	}

	maxPreLength := len(strs[0])

	for i := 0; i < strlen; i++ {
		if len(strs[i]) < maxPreLength {
			maxPreLength = len(strs[i])
		}
	}

	i := 0
Cir:
	for ; i < maxPreLength; i++ {
		for j := 1; j < strlen; j++ {
			if strs[j][i] != strs[j-1][i] {
				break Cir
			}
		}
	}
	return strs[0][0:i]
}
