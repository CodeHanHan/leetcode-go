package topic006

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	b := numRows - 1
	strArr := make([]string, numRows)
	for k, v := range s {
		if (k/b)%2 == 0 { // 偶数组, 正序赋值
			strArr[k%b] += string(v)
		} else { // 奇数组, 倒序赋值
			strArr[b-(k%b)] += string(v)
		}
	}
	return strings.Join(strArr, "")
}

func convert_1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	strArr := make([]string, numRows)
	down := false
	curRow := 0

	for _, v := range s {
		strArr[curRow] += string(v)
		if curRow == 0 || curRow == numRows-1 {
			down = !down
		}

		if down {
			curRow++
		} else {
			curRow--
		}
	}

	return strings.Join(strArr, "")
}
