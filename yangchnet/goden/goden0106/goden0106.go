package goden0106

import (
	"fmt"
	"strings"
)

func compressString(S string) string {
	if len(S) <= 1 {
		return S
	}

	str := strings.Builder{}
	str.WriteByte(S[0])
	i, j := 1, 1

	curByte := S[0]
	var count int = 1
	for i < len(S) {
		if S[i] == curByte {
			i++
			count += 1
		} else {
			curByte = S[i]
			str.WriteString(fmt.Sprint(count))
			str.WriteByte(curByte)
			j += 2
			i++
			count = 1
		}
	}
	str.WriteString(fmt.Sprint(count))

	if len(str.String()) >= len(S) {
		return S
	}

	return str.String()
}
