package lc394

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	numStack := []int{}
	strStack := []string{}
	num := 0
	character := ""
	for _, char := range s {
		if char >= '0' && char <= '9' {
			n, _ := strconv.Atoi(string(char))
			num = num*10 + n
		} else if char == '[' {
			strStack = append(strStack, character)
			character = ""
			numStack = append(numStack, num)
			num = 0
		} else if char == ']' {
			count := numStack[len(numStack)-1]
			numStack = numStack[:len(numStack)-1]
			str := strStack[len(strStack)-1]
			strStack = strStack[:len(strStack)-1]
			character = string(str) + strings.Repeat(character, count)
		} else {
			character += string(char)
		}
	}
	return character
}
