package offer_65_s

func add(a int, b int) int {
	// 进位
	var carry int
	for b != 0 {
		// 进位
		carry = (a & b) << 1
		// 不进位加
		a ^= b
		// 加进位
		b = carry
	}
	return a
}
