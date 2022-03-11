package dynamic

import "github.com/CodeHanHan/leetcode-go/algo"

/*
莱文斯坦距离指，只用增加、删除、替换字符这三个编辑操作，从一个字符串变更到另一个字符串的操作次数

莱文斯坦距离的大小，表示两个字符串差异的大小；
*/

/** 状态转移方程
如果：a[i]!=b[j]，那么：min_edist(i, j)就等于：
min(min_edist(i-1,j)+1, min_edist(i,j-1)+1, min_edist(i-1,j-1)+1)

如果：a[i]==b[j]，那么：min_edist(i, j)就等于：
min(min_edist(i-1,j)+1, min_edist(i,j-1)+1，min_edist(i-1,j-1))

其中，min表示求三数中的最小值。
*/

func LevenshteinDist(str1, str2 string) int {
	dp := make([][]int, len(str1))
	for i := range dp {
		dp[i] = make([]int, len(str2))
	}

	for column := 0; column < len(str2); column++ {
		if str1[0] == str2[column] {
			// 如果str2的当前字符和str1的第一个字符相等，相当于只要删除前[column]个字符二者就可相同
			dp[0][column] = column
		} else if column != 0 {
			// 如果str2的当前字符和str1的第一个字符不等, 需要从前一个状态进行一次编辑操作
			dp[0][column] = dp[0][column-1] + 1
		} else {
			// str[2] 与 str[1]不同
			dp[0][column] = 1
		}
	}

	for row := 1; row < len(str1); row++ {
		if str1[row] == str2[0] {
			dp[row][0] = row
		} else if row != 0 {
			dp[row][0] = dp[row-1][0] + 1
		} else {
			dp[row][0] = 1
		}
	}

	for row := 1; row < len(str1); row++ {
		for column := 1; column < len(str2); column++ {
			if str1[row] == str2[column] {
				dp[row][column] = algo.MinN(
					dp[row-1][column]+1,
					dp[row][column-1]+1,
					dp[row-1][column-1])
			} else {
				dp[row][column] = algo.MinN(
					dp[row-1][column]+1,
					dp[row][column-1]+1,
					dp[row-1][column-1]+1)
			}
		}
	}

	return dp[len(str1)-1][len(str2)-1]
}
