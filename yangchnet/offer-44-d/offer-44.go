package offer_44_d_

import "strconv"

func findNthDigit(n int) int {
	digit := 1
	start := 1
	count := 9

	for n > count {
		n -= count
		start *= 10
		digit += 1
		count = 9 * start * digit
	}
	num := start + (n-1)/digit
	return int(strconv.Itoa(num)[(n-1)%digit] - '0')
}
