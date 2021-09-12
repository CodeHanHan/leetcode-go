package offer_10_I_s

func fib(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	b := 1
	for i := 1; i < n; i++ {
		a, b = b, (a+b)%1000000007
	}
	return b % 1000000007
}
