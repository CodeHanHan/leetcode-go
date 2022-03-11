package dynamic

/*
一只青蛙一次可以跳上1级台阶，也可以跳上2级台阶。求该青蛙跳上一个 n级的台阶总共有多少种跳法。
*/

// 状态转移方程： f(n) = f(n-1) + f(n-2)

func FrogSteps(n int) int {
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
