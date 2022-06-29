package topic97

// 超出时间限制
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) == 0 && len(s2) == 0 && len(s3) == 0 {
		return true
	}
	if len(s3) == 0 && (len(s1) != 0 || len(s2) != 0) {
		return false
	}

	i := 0
	for len(s1) > i && len(s3) > i && s1[i] == s3[i] {
		i++
	}

	j := 0
	for len(s2) > j && len(s3) > j && s2[j] == s3[j] {
		j++
	}

	var x, y bool = false, false

	for ii := 1; ii <= i; ii++ {
		x = x || isInterleave(s1[ii:], s2, s3[ii:])
	}

	for jj := 1; jj <= j; jj++ {
		y = y || isInterleave(s1, s2[jj:], s3[jj:])
	}

	return x || y
}

func isInterleave1(s1 string, s2 string, s3 string) bool {
	n, m, t := len(s1), len(s2), len(s3)
	if (n + m) != t {
		return false
	}
	f := make([][]bool, n+1) // f[i][j] 表示s1的前i个元素和s2的前j个元素是否能交错组成s3的前i+j个元素
	for i := 0; i <= n; i++ {
		f[i] = make([]bool, m+1)
	}
	f[0][0] = true
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			p := i + j - 1
			if i > 0 {
				f[i][j] = f[i][j] || (f[i-1][j] && s1[i-1] == s3[p])
			}
			if j > 0 {
				f[i][j] = f[i][j] || (f[i][j-1] && s2[j-1] == s3[p])
			}
		}
	}
	return f[n][m]
}
