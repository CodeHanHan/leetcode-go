package offer_58_II_s

import "strings"

func reverseLeftWords(s string, n int) string {
	sl := strings.Split(s, "")
	return strings.Join(append(sl[n:], sl[:n]...), "")
}
