package topic58

func lengthOfLastWord(s string) int {
	n := len(s)
	i := n - 1
	for i >= 0 && s[i] == ' ' {
		i--
	}

	end := i

	for i >= 0 && s[i] != ' ' {
		i--
	}

	return end - i
}
