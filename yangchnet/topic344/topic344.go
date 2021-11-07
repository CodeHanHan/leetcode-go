package topic344

func reverseString(s []byte) {
	var i, j int = 0, len(s) - 1

	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
