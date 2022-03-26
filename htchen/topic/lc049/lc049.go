package lc049

import "sort"

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	keyMap := make(map[string][]string, 0)

	for _, str := range strs {
		key := getKey(str)
		keyMap[key] = append(keyMap[key], str)
	}
	for _, value := range keyMap {
		result = append(result, value)
	}
	return result
}

func getKey(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool { return bytes[i] < bytes[j] })
	return string(bytes)
}
