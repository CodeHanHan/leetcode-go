package goden0102

import "reflect"

// 解法2
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m1 := make(map[rune]int)
	m2 := make(map[rune]int)

	for _, v := range s1 {
		m1[v]++
	}

	for _, v := range s2 {
		m2[v]++
	}

	return reflect.DeepEqual(m1, m2)
}
