package LC009

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	tempx := x
	res := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		res = res*10 + pop
	}
	//fmt.Println(res)
	return res == tempx
}
