package topic49

import "sort"

func groupAnagrams(strs []string) [][]string {
	n := len(strs)
	if n <= 1 {
		return [][]string{strs}
	}

	ans := make([][]string, 0)

	mm := make(map[string][]string)
	for _, v := range strs {
		v1 := append([]byte{}, append([]byte(nil), v[:]...)...)
		sort.Slice(v1[:], func(i, j int) bool {
			return v1[i] < v1[j]
		})

		if _, ok := mm[string(v1)]; !ok {
			s := make([]string, 0)
			s = append(s, v)
			mm[string(v1)] = s
		} else {
			mm[string(v1)] = append(mm[string(v1)], v)
		}

	}

	for _, v := range mm {
		ans = append(ans, v)
	}

	return ans
}
