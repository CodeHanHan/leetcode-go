package goden0109

func isFlipedString(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if s1 == s2 {
		return true
	}

	n := len(s1)
	h := 0
	for h < n {
		if s1[0:n-h] == s2[h:] && s1[n-h:] == s2[0:h] {
			return true
		}
		h++
	}

	return false
}
