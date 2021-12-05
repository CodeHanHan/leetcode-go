package lc171

func titleToNumber(columnTitle string) int {
	n := len(columnTitle)
	number := 0
	multiple := 1
	for i := n - 1; i >= 0; i-- {
		k := columnTitle[i] - 'A' + 1
		number += int(k) * multiple
		multiple *= 26
	}
	return number
}
