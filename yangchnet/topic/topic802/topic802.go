package topic802

/*
白色（用 0 表示）：该节点尚未被访问；
灰色（用 1 表示）：该节点位于递归栈中，或者在某个环上；
黑色（用 2 表示）：该节点搜索完毕，是一个安全节点。
*/

func eventualSafeNodes(graph [][]int) []int {
	color := make([]int, len(graph))

	var safty func(start int) bool
	safty = func(start int) bool {
		if color[start] > 0 {
			return color[start] == 2
		}

		color[start] = 1

		for i := 0; i < len(graph[start]); i++ {
			if !safty(graph[start][i]) {
				return false
			}
		}

		color[start] = 2

		return true
	}

	for i := 0; i < len(graph); i++ {
		safty(i)
	}

	ans := make([]int, 0)
	for i := 0; i < len(graph); i++ {
		if color[i] == 2 {
			ans = append(ans, i)
		}
	}

	return ans
}
