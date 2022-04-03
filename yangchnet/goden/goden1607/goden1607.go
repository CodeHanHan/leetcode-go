package goden1607

import "unsafe"

func maximum(a int, b int) int {
	size := int(unsafe.Sizeof(a))
	k := (b - a) >> ((size * 8) - 1) & 0x1

	return a*k + b*(k^1)
}
