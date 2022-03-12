package goden0802

// 回溯法-未通过
func pathWithObstacles(obstacleGrid [][]int) [][]int {
	var res [][]int = make([][]int, 0)
	visited := make([][]bool, len(obstacleGrid))
	for i := range visited {
		visited[i] = make([]bool, len(obstacleGrid[0]))
	}

	var fn func(grid [][]int, path [][]int, i, j int)
	fn = func(grid [][]int, path [][]int, i, j int) {
		if visited[i][j] {
			return
		}

		visited[i][j] = true
		if i == 0 && j == 0 && grid[i][j] >= 1 {
			return
		}

		if i == len(grid)-1 && j == len(grid[0])-1 && grid[i][j] <= 0 {
			path = append(path, []int{i, j})
			res = path
			return
		}

		if j+1 < len(grid[0]) && grid[i][j+1] <= 0 { //向右走
			path = append(path, []int{i, j})
			fn(grid, path, i, j+1)
			path = path[:len(path)-1]
		}

		if i+1 < len(grid) && grid[i+1][j] <= 0 { // 向下走
			path = append(path, []int{i, j})
			fn(grid, path, i+1, j)
			path = path[:len(path)-1]
		}

	}

	fn(obstacleGrid, [][]int{}, 0, 0)

	return res
}

/*
func pathWithObstacles(obstacleGrid [][]int) [][]int {
    var (
        ans [][]int
        r int = len(obstacleGrid) - 1
        c int = len(obstacleGrid[0]) - 1
        f func([][]int)
    )
    f = func(path [][]int) {
        if len(ans) == 0 {
            i, j := path[len(path) - 1][0], path[len(path) - 1][1]
            if obstacleGrid[i][j] == 0 {
                obstacleGrid[i][j] = 1
                if j < c {
                    f(append(path, []int{i, j + 1}))
                }
                if i < r {
                    f(append(path, []int{i + 1, j}))
                }
                if i == r && j == c{
                    ans = make([][]int, r + c + 1)
                    copy(ans, path)
                }
            }
        }
    }
    f([][]int{{0, 0}})
    return ans
}
*/
