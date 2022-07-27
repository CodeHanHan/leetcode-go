package topic415

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	res := ""
	carry := 0
	i, j := len(num1)-1, len(num2)-1
	for i >= 0 || j >= 0 || carry != 0 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		add := x + y + carry
		res = fmt.Sprintf("%d", add%10) + res
		carry = add / 10
		i--
		j--
	}
	return res
}
