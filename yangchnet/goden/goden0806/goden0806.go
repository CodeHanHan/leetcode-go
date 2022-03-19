package goden0806

func hanota(A []int, B []int, C []int) []int {
	if len(A) <= 0 {
		return C
	}

	// fn: 将A上面n个通过B移动到C
	var fn func(A, B, C *[]int, n int)
	fn = func(A, B, C *[]int, n int) {
		if n == 1 {
			*C = append((*C), (*A)[len(*A)-1])
			*A = (*A)[:len(*A)-1]
			return
		}

		// 通过C将A上n-1个移动到B
		fn(A, C, B, n-1)
		*C = append((*C), (*A)[len(*A)-1])
		*A = (*A)[:len(*A)-1]
		// 通过A，将B上n-1个移动到C上
		fn(B, A, C, n-1)
	}

	fn(&A, &B, &C, len(A))

	return C
}
