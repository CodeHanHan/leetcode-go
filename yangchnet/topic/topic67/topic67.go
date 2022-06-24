package topic67

import (
	"fmt"
	"strings"
)

func addBinary(a string, b string) string {
	ans := make([]string, 0)
	m := map[byte]int{
		'1': 1,
		'0': 0,
	}

	var c int = 0
	i := 0
	for ; i < min(len(a), len(b)); i++ {
		x := m[a[len(a)-i-1]]
		y := m[b[len(b)-i-1]]
		ans = append([]string{fmt.Sprintf("%d", (x+y+c)%2)}, ans...)
		c = (x + y + c) / 2
	}

	for ; i < len(a); i++ {
		x := m[a[len(a)-i-1]]
		ans = append([]string{fmt.Sprintf("%d", (x+c)%2)}, ans...)
		c = (x + c) / 2
	}

	for ; i < len(b); i++ {
		y := m[b[len(b)-i-1]]
		ans = append([]string{fmt.Sprintf("%d", (y+c)%2)}, ans...)
		c = (y + c) / 2
	}

	if c != 0 {
		ans = append([]string{fmt.Sprintf("%d", c)}, ans...)
	}

	return strings.Join(ans, "")
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
