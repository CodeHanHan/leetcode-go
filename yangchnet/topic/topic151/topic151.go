package topic151

import "strings"

func reverseWords(s string) string {
	list := strings.Split(s, " ")

	ret := make([]string, 0)

	for i := len(list) - 1; i >= 0; i-- {
		if list[i] != "" {
			ret = append(ret, list[i])
		}
	}

	return strings.Join(ret, " ")
}

func reverseWords1(s string) string {
	for len(s) >= 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}

	for len(s) >= 0 && s[0] == ' ' {
		s = s[1:]
	}

	n := strings.Index(s, " ")

	if n == -1 {
		return s
	}

	return reverseWords1(s[n+1:]) + " " + s[0:n]
}
