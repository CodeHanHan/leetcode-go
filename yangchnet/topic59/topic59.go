package topic59

type direction [2]int

var (
	north direction = [2]int{-1, 0} // ↑
	south direction = [2]int{1, 0}  // ↓
	west  direction = [2]int{0, -1} // ←
	east  direction = [2]int{0, 1}  // →
)

var drcs [4]direction = [4]direction{east, south, west, north}

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
	}

	var x, y int = 0, -1
	var curDrc int = 0
	for i := 1; i <= n*n; i++ {
		p, q := step(x, y, curDrc)
		if p < 0 || p >= n || q < 0 || q >= n || matrix[p][q] > 0 {
			x, y, curDrc = switchDrc(x, y, curDrc)
			matrix[x][y] = i
			continue
		} else {
			x, y = p, q
			matrix[x][y] = i
		}
	}

	return matrix
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
