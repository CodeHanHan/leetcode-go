package topic54

type direction [2]int

var (
	north direction = [2]int{-1, 0} // ↑
	south direction = [2]int{1, 0}  // ↓
	west  direction = [2]int{0, -1} // ←
	east  direction = [2]int{0, 1}  // →
)

var drcs [4]direction = [4]direction{east, south, west, north}

func spiralOrder(matrix [][]int) []int {
	nums := make([]int, 0)
	visited := make([][]bool, len(matrix))
	for i, _ := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}

	var i, j int = 0, -1 // 设置起始点

	var curDrc int = 0 // 设置初始方向

	for {
		x, y := step(i, j, curDrc)                                                       // x, y 是下一步要经过的坐标
		if x >= 0 && x < len(matrix) && y >= 0 && y < len(matrix[0]) && !visited[x][y] { // x, y 验证未越界，未访问过，则下脚
			i, j = x, y
		} else { // x, y 验证越界或已被访问过，则改变方向
			i, j, curDrc = switchDrc(i, j, curDrc)
			if i < 0 || i >= len(matrix) || j < 0 || j >= len(matrix[0]) || visited[i][j] { // 改变方向后仍已被访问过或是越界，则说明遍历完毕，返回
				break
			}
		}
		visited[i][j] = true
		nums = append(nums, matrix[i][j])

	}
	return nums
}

func step(i, j int, drc int) (x int, y int) {
	x = i + drcs[drc][0]
	y = j + drcs[drc][1]
	return
}

func switchDrc(i, j int, curDrc int) (x int, y int, drc int) {
	drc = (curDrc + 1) % 4
	x = i + drcs[drc][0]
	y = j + drcs[drc][1]
	return
}

func spiralOrder1(matrix [][]int) []int {
	ans, i, j, di, dj := []int{}, 0, 0, 0, 1
	if len(matrix) < 1 {
		return ans
	}

	n, m := len(matrix), len(matrix[0])

	for k := 0; k < n*m; k++ {
		ans = append(ans, matrix[i][j])
		matrix[i][j] = 0
		if matrix[(i+di+n)%len(matrix)][(j+dj+m)%len(matrix[0])] == 0 {
			di, dj = dj, -di
		}

		i += di
		j += dj
	}

	return ans
}
