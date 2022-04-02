package goden1002

import (
	"sort"
)

// 排序
func groupAnagrams1(strs []string) [][]string {
	n := len(strs)

	_strs := make([]string, 0)

	for k := 0; k < n; k++ {
		b := []byte(strs[k])
		sort.Slice(b, func(i int, j int) bool {
			return b[i] < b[j]
		})
		_strs = append(_strs, string(b))
	}

	m := make(map[string][]string)

	for i := 0; i < n; i++ {
		if _, ok := m[_strs[i]]; !ok {
			m[_strs[i]] = make([]string, 0)
		}

		m[_strs[i]] = append(m[_strs[i]], strs[i])
	}

	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}

// 质数求积
func groupAnagrams2(strs []string) [][]string {
	q := [26]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	m := make(map[int][]string)
	for i := 0; i < len(strs); i++ {
		mul := 1
		for j := 0; j < len(strs[i]); j++ {
			mul = mul * q[int(strs[i][j]-'a')]
		}

		if _, ok := m[mul]; !ok {
			m[mul] = make([]string, 0)
		}

		m[mul] = append(m[mul], strs[i])
	}

	res := make([][]string, 0)
	for _, v := range m {
		res = append(res, v)
	}

	return res
}
