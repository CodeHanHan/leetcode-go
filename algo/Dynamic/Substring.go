package dynamic

import "github.com/CodeHanHan/leetcode-go/algo"

/*
最长公共子串
表征的也是两个字符串之间的相似程度
*/

/** 状态转移方程
如果：a[i]==b[j]，那么：max_lcs(i, j)就等于：
max(max_lcs(i-1,j-1)+1, max_lcs(i-1, j), max_lcs(i, j-1))；

如果：a[i]!=b[j]，那么：max_lcs(i, j)就等于：
max(max_lcs(i-1,j-1), max_lcs(i-1, j), max_lcs(i, j-1))；
*/

func LongestCommonSubString(s1, s2 string) int {
	dp := make([][]int, len(s1))
	for i := range dp {
		dp[i] = make([]int, len(s2))
	}

	// 初始化第0列
	for row := 0; row < len(s1); row++ {
		if s1[row] == s2[0] { // 字符相等
			dp[row][0] = 1
		} else if row != 0 { // s1与s2当前字符不相等
			dp[row][0] = dp[row-1][0]
		} else { // s1第一个字符与是s2不相等
			dp[row][0] = 0
		}
	}

	// 初始化第0行
	for column := 1; column < len(s2); column++ {
		if s2[column] == s1[0] {
			dp[0][column] = 1
		} else if column != 0 {
			dp[0][column] = dp[0][column-1]
		} else {
			dp[0][column] = 0
		}
	}

	for row := 1; row < len(s1); row++ {
		for column := 1; column < len(s2); column++ {
			if s1[row] == s2[column] {
				dp[row][column] = algo.MaxN(
					dp[row-1][column],
					dp[row][column-1],
					dp[row-1][column-1]+1,
				)
			} else {
				dp[row][column] = algo.MaxN(
					dp[row-1][column],
					dp[row][column-1],
					dp[row-1][column-1],
				)
			}
		}
	}

	return dp[len(s1)-1][len(s2)-1]
}
