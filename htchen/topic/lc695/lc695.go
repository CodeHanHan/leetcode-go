package lc695

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxArea, curArea := 0, 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x >= m || x < 0 || y >= n || y < 0 {
			return
		}
		if grid[x][y] == 0 {
			return
		}
		grid[x][y] = 0
		curArea++
		for _, d := range dir4 {
			dfs(x+d.x, y+d.y)
		}
	}

	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				curArea = 0
				dfs(i, j)
			}
			maxArea = max(maxArea, curArea)
		}
	}
	return maxArea
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
