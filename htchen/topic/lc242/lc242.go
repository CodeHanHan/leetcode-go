package lc242

import "sort"

func isAnagram1(s string, t string) bool {
	s1, t1 := []byte(s), []byte(t)
	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
	sort.Slice(t1, func(i, j int) bool { return t1[i] < t1[j] })
	return string(s1) == string(t1)
}

func isAnagram2(s string, t string) bool {
	var c1, c2 [26]int
	for _, ch := range s {
		c1[ch-'a']++
	}
	for _, ch := range t {
		c2[ch-'a']++
	}
	return c1 == c2
}
