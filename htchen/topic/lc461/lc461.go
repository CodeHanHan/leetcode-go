package lc461

import "math/bits"

func hammingDistance1(x int, y int) int {
	var res int
	for s := x ^ y; s > 0; s >>= 1 {
		res += s & 1
	}
	return res
}

func hammingDistance2(x int, y int) int {
	return bits.OnesCount(uint(x ^ y))
}
