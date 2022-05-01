package lc338

func countBits(n int) []int {
	bits := make([]int, n+1)
	for i, _ := range bits {
		bits[i] = onesCount(i)
	}
	return bits
}

func onesCount(x int) int {
	cnt := 0
	for ; x > 0; x &= x - 1 {
		cnt++
	}
	return cnt
}
