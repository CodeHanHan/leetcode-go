package goden0501

import "fmt"

func insertBits(N int, M int, i int, j int) int {
	left := N >> (uint(j) + 1) << (uint(j) + 1)
	a := N >> uint(i) << uint(i)
	fmt.Println(a)
	right := (N >> uint(i) << uint(i)) ^ N
	lr := left | right
	return lr | M<<uint(i)
}
