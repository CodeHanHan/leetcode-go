package LC058

func lengthOfLastWord(s string) int {
	index := len(s) - 1
	var length int

	for s[index] == ' ' {
		index--
	}

	for index >= 0 && s[index] != ' ' {
		length++
		index--
	}
	return length
}
