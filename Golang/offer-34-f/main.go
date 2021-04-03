package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type rou struct {
	route  []*TreeNode // 路径
	array  []int
	routed bool // 是否路由到此
}

func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	var dp = make([][]rou, sum+1, sum+1)
	for i := 0; i < sum+1; i++ {
		dp[i] = make([]rou, sum+1, sum+1)
	}

	dp[0][root.Val].route = []*TreeNode{root}
	dp[0][root.Val].routed = true
	dp[0][root.Val].array = append([]int{}, root.Val)
	for i := 0; i <= sum; i++ { // 行
		for j := 0; j <= sum; j++ { //列
			if dp[i][j].routed { //这个累加数字可以达到
				left := dp[i][j].route[len(dp[i][j].route)-1].Left
				right := dp[i][j].route[len(dp[i][j].route)-1].Right

				if left != nil && j+left.Val <= sum {
					dp[i+1][j+left.Val].routed = true
					dp[i+1][j+left.Val].route = append(dp[i][j].route, left)
					dp[i+1][j+left.Val].array = append(dp[i][j].array, left.Val)
					if isleaf(dp[i+1][j+left.Val]) && j+left.Val == sum {
						var a []int = make([]int, len(dp[i+1][j+left.Val].array), len(dp[i+1][j+left.Val].array))
						copy(a, dp[i+1][j+left.Val].array)
						//a := dp[i+1][j+left.Val].array
						res = append(res, a)
					}

				}
				if right != nil && j+right.Val <= sum {
					dp[i+1][j+right.Val].routed = true
					dp[i+1][j+right.Val].route = append(dp[i][j].route, right)
					dp[i+1][j+right.Val].array = append(dp[i][j].array, right.Val)
					if isleaf(dp[i+1][j+right.Val]) && j+right.Val == sum {
						var b []int = make([]int, len(dp[i+1][j+right.Val].array), len(dp[i+1][j+right.Val].array))
						copy(b, dp[i+1][j+right.Val].array)
						//b := dp[i+1][j+right.Val].array
						res = append(res, b)
					}
				}
			}
		}
	}
	return res
}

func isleaf(r rou) bool {
	left := r.route[len(r.route)-1].Left
	right := r.route[len(r.route)-1].Right
	if left == nil && right == nil {
		return true
	}
	return false
}
