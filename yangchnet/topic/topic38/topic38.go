package topic38

import (
	"bytes"
	"fmt"
)

func countAndSay(n int) string {
	des := make([]string, n+1)

	des[1] = "1"
	for i := 2; i <= n; i++ {
		des[i] = desPre(des[i-1])
	}

	return des[n]
}

func desPre(in string) string {
	buf := bytes.Buffer{}
	for i := 0; i < len(in); {
		count := 1
		cur := in[i]

		j := i + 1
		for ; j < len(in) && in[j] == cur; j++ {
			count++
		}

		buf.WriteString(fmt.Sprintf("%d%d", count, cur-'1'+1))
		i = j

	}

	return buf.String()
}
