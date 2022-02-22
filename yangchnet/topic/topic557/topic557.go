package topic557

import "bytes"

func reverseWords(s string) string {
	var buf bytes.Buffer

	for i, v := range s {
		if v == ' ' {
			for j := i - 1; j >= 0 && s[j] != ' '; j-- {
				buf.WriteByte(s[j])
			}
			buf.WriteByte(' ')
		}
	}

	for j := len(s) - 1; j >= 0 && s[j] != ' '; j-- {
		buf.WriteByte(s[j])
	}

	return buf.String()
}
