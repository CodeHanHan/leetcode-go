package goden1715

import "sort"

func longestWord(words []string) string {
	existed := make(map[string]bool)
	for _, word := range words {
		existed[word] = true
	}

	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) < len(words[j]) {
			return false
		} else if len(words[i]) > len(words[j]) {
			return true
		}

		return words[i] < words[j]
	})

	for _, word := range words {
		existed[word] = false
		if check(existed, word) {
			return word
		}
		existed[word] = true
	}

	return ""
}

func check(existed map[string]bool, word string) bool {
	if word == "" || existed[word] {
		return true
	}

	for i := 0; i < len(word); i++ {
		if existed[word[:i]] && check(existed, word[i:]) {
			return true
		}
	}

	return false
}
