package lc463

type pair struct{ x, y int }

var dir4 = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

//可统计存在多个岛屿的情况
func islandPerimeter(grid [][]int) int {
	row, column := len(grid), len(grid[0])
	girth := 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || x >= row || y >= column || y < 0 || grid[x][y] == 0 {
			girth++
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
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
			}
		}
	}
	return girth
}

//统计题目中仅存在一个岛屿的情况
func islandPerimeter1(grid [][]int) int {
	row, column := len(grid), len(grid[0])
	girth := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == 1 {
				for _, d := range dir4 {
					if x, y := i+d.x, j+d.y; x < 0 || x >= row || y < 0 || y >= column || grid[x][y] == 0 {
						girth++
					}
				}
			}
		}
	}
	return girth
}
