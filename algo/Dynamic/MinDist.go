package dynamic

import "github.com/CodeHanHan/leetcode-go/algo"

/*
有一个 n 乘以 n 的矩阵 w[n][n]。
矩阵存储的都是正整数。棋子起始位置在左上角，终止位置在右下角。
将棋子从左上角移动到右下角。每次只能向右或者向下移动一位。
从左上角到右下角，会有很多不同的路径可以走。把每条路径经过的数字加起来看作路径的长度。
那从左上角移动到右下角的最短路径长度是多少呢
*/

// 状态转移方程： minDist[i][j] = matrix[i][j] + min(minDist[i-1][j], minDist[i][j-1])

func MinDist(matrix [][]int) int {
	// 构造矩阵
	pm := make([][]int, len(matrix))
	for i := range pm {
		pm[i] = make([]int, len(matrix[0]))
	}

	// 初始化
	pm[0][0] = matrix[0][0]
	for column := 1; column < len(pm[0]); column++ { // 行初始化
		pm[0][column] = matrix[0][column] + pm[0][column-1]
	}
	for row := 1; row < len(pm); row++ { // 列初始化
		pm[row][0] = matrix[row][0] + pm[row-1][0]
	}

	for i := 1; i < len(pm); i++ {
		for j := 1; j < len(pm[0]); j++ {
			pm[i][j] = matrix[i][j] + algo.MinN(pm[i-1][j], pm[i][j-1])
		}
	}

	return pm[len(pm)-1][len(pm[0])-1]
}
