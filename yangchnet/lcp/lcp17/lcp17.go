package lcp17

func calculate(s string) int {
	x, y := 1, 0

	for _, v := range s {
		if v == 'A' {
			x = 2*x + y
		} else if v == 'B' {
			y = 2*y + x
		}
	}

	return x + y
}
