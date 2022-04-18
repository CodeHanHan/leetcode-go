package goden1724

import "math"

func getMaxMatrix(matrix [][]int) []int {
	ans := make([]int, 4) //保存最大子矩阵的左上角和右下角的行列坐标
	n := len(matrix)
	m := len(matrix[0])

	sum := 0                //相当于dp[i],dp_i
	maxSum := math.MinInt64 //记录最大值
	bestr1, bestc1 := 0, 0  //暂时记录左上角，相当于begin

	for i := 0; i < n; i++ { //以i为上边，从上而下扫描
		b := make([]int, m)      //记录当前i~j行组成大矩阵的每一列的和，将二维转化为一维;每次更换子矩形上边，就要清空b，重新计算每列的和
		for j := i; j < n; j++ { //子矩阵的下边，从i到N-1，不断增加子矩阵的高
			//一下就相当于求一次最大子序列和
			sum = 0 //从头开始求dp
			for k := 0; k < m; k++ {
				b[k] += matrix[j][k]
				//我们只是不断增加其高，也就是下移矩阵下边，所有这个矩阵每列的和只需要加上新加的哪一行的元素
				//因为我们求dp[i]的时候只需要dp[i-1]和nums[i],所有在我们不断更新b数组时就可以求出当前位置的dp_i
				if sum > 0 {
					sum += b[k]
				} else {
					sum = b[k] // 暂时保存其左上角
					bestr1 = i
					bestc1 = k
				}
				if sum > maxSum {
					maxSum = sum
					ans[0] = bestr1 //更新答案
					ans[1] = bestc1
					ans[2] = j
					ans[3] = k
				}
			}
		}
	}

	return ans
}
