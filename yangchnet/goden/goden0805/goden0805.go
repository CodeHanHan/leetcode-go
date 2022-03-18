package goden0805

func multiply(A int, B int) int {
	if B == 0 {
		return 0
	}

	return A + multiply(A, B-1)
}
