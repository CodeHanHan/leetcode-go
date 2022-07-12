package lc1254

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func closedIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cnt := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x >= m || x < 0 || y >= n || y < 0 {
			return
		}
		if grid[x][y] == 1 {
			return
		}
		if grid[x][y] == 2 {
			return
		}
		grid[x][y] = 2
		for _, d := range dir4 {
			dfs(x+d.x, y+d.y)
		}
	}

	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}

	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] == 0 {
				cnt++
				dfs(i, j)
			}
		}
	}
	return cnt
}
