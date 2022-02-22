package topic028

func strStr(haystack string, needle string) int {
	if len(needle) <= 0 {
		return 0
	}

	if len(needle) > len(haystack) {
		return -1
	}

	var i int = 0
	for i < len(haystack) {
		if haystack[i] == needle[0] && i+len(needle) <= len(haystack) && strEqual(haystack[i:i+len(needle)], needle) { // 第一个字符匹配成功
			return i
		}
		i++
	}

	return -1
}

func strEqual(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
