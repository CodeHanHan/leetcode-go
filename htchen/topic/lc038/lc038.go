package lc038

import "strconv"

func countAndSay(n int) string {
	content := "1"
	for n > 1 {
		content = say(content)
		n--
	}
	return content
}

func say(content string) string {
	content = content + " "
	res := ""
	lens := 0
	left := 0
	for right := 1; right < len(content); {
		if content[right] != content[left] {
			lens = right - left
			res = res + strconv.Itoa(lens) + string(content[left])
			left = right
		}
		right++
	}
	return res
}
