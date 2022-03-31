package goden1001

func merge(A []int, m int, B []int, n int) {
	i, j := m-1, n-1
	cur := len(A) - 1

	for i >= 0 || j >= 0 {
		if i == -1 {
			A[cur] = B[j]
			j--
		} else if j == -1 {
			A[cur] = A[i]
			i--
		} else if A[i] > B[j] {
			A[cur] = A[i]
			i--
		} else {
			A[cur] = B[j]
			j--
		}

		cur--
	}
}
