package LC028

func strStr(haystack string, needle string) int {
	if len(haystack) == 0 || len(haystack) < len(needle) {
		return -1
	}
	if len(needle) == 0 {
		return 0
	}

	mlen := len(haystack)
	slen := len(needle)
	for i := 0; i <= mlen-slen; i++ {
		if haystack[i:i+slen] == needle {
			return i
		}
	}
	return -1
}
