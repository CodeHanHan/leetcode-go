package offer_43_d_

func countDigitOne(n int) int {
	high := n / 10
	cur := n % 10
	low := 0
	digit := 1

	sum1 := 0

	for high != 0 || cur != 0 {
		switch cur {
		case 0:
			s := high * digit
			sum1 += s
		case 1:
			s := high*digit + low + 1
			sum1 += s
		default:
			s := (high + 1) * digit
			sum1 += s
		}
		low += cur * digit
		cur = high % 10
		high = high / 10
		digit *= 10
	}
	return sum1
}
