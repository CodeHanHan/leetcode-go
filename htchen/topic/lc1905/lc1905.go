package lc1905

var ok bool

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	cnt := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		row, column := len(grid2), len(grid2[0])
		if r >= row || r < 0 || c >= column || c < 0 {
			return
		}
		if grid2[r][c] == 0 {
			return
		}
		grid2[r][c] = 0
		if grid1[r][c] == 0 {
			ok = false
		}
		dfs(r-1, c)
		dfs(r+1, c)
		dfs(r, c-1)
		dfs(r, c+1)
	}

	for i, row := range grid2 {
		for j, v := range row {
			if v > 0 {
				ok = true
				dfs(i, j)
				if ok == true {
					cnt++
				}
			}
		}
	}
	return cnt
}
