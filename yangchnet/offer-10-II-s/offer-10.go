package offer_10_II_s

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	a := 1
	b := 2
	for i := 2; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b
}
