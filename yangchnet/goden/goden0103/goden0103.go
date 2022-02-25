package goden0103

import "strings"

func replaceSpaces(S string, length int) string {
	// var buf bytes.Buffer
	// count := 0
	// for _, s := range S {
	// 	if s == ' ' {
	// 		buf.WriteString("%20")
	// 		count++
	// 	} else {
	// 		buf.WriteRune(s)
	// 		count++
	// 	}
	// 	if count >= length {
	// 		break
	// 	}
	// }

	// return buf.String()

	return strings.Replace(S[0:length], " ", "%20", -1)
}
