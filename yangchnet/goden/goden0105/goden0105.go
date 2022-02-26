package goden0105

import "math"

func oneEditAway(first string, second string) bool {
	if first == second {
		return true
	}
	if int(math.Abs(float64(len(first)-len(second)))) > 1 {
		return false
	}

	if max(len(first), len(second)) <= 1 {
		return true
	}

	if first[0] == second[0] {
		return oneEditAway(first[1:], second[1:])
	} else if first[len(first)-1] == second[len(second)-1] {
		return oneEditAway(first[:len(first)-1], second[:len(second)-1])
	} else {
		return false
	}

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func oneEditAway2(first string, second string) bool {
	if first == second {
		return true
	}
	M, N := len(first), len(second)
	// 长在前  即令first始终是长的，M也表示first的长度，方便解题
	if M < N {
		first, second = second, first
		M, N = N, M
	}
	// 如果长度差值大于1 则直接false
	if M-N > 1 {
		return false
	}
	// 核心代码
	i := 0
	// 遍历较短的   遇到不同的字符，比较 （first删除该位置 || 两者替换该位置）  的情况下后面是否相等
	for i < N {
		if first[i] != second[i] {
			return first[i+1:] == second[i:] || first[i+1:] == second[i+1:]
		}
		i++
	}
	return true
}
