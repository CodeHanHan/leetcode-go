package topic11

import (
	"bytes"
)

// 纵向扫描法
func longestCommonPrefix(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}

	maxPrefix := bytes.Buffer{}

	for k := 0; k < len(strs[0]); k++ { // k代表检查到第k个字符
		for p := 0; p < len(strs); p++ { // p代表检查到第p个数组元素
			if k > len(strs[p])-1 {
				return maxPrefix.String()
			}
			if strs[p][k] != strs[0][k] {
				return maxPrefix.String()
			}
		}
		maxPrefix.WriteByte(strs[0][k])
	}
	return maxPrefix.String()
}

// 横向扫描法
func longestCommonPrefix_1(strs []string) string {
	if len(strs) <= 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	var maxPrefix string = commonPrefix(strs[0], strs[1])
	for i := 1; i < len(strs); i++ {
		maxPrefix = commonPrefix(maxPrefix, strs[i])
	}
	return maxPrefix
}

func commonPrefix(str1, str2 string) string {
	prefix := bytes.Buffer{}

	for i := 0; i < len(str1); i++ {
		if i > len(str2)-1 {
			return prefix.String()
		}
		if str1[i] != str2[i] {
			return prefix.String()
		}
		prefix.WriteByte(str1[i])
	}
	return prefix.String()
}
