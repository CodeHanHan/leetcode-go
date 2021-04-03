package offer_58_I_s

import "strings"

func reverseWords(s string) string {
	res := make([]string, 0)
	split := strings.Split(s, " ")
	for i := len(split) - 1; i >= 0; i-- {
		if split[i] != "" {
			res = append(res, split[i])
		}
	}
	return strings.Join(res, " ")
}
