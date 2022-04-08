package goden1622

func printKMoves(K int) []string {
	direction := [4]byte{'R', 'D', 'L', 'U'}
	color := [2]byte{'_', 'X'}

	board := make(map[[2]int]int)

	board[[2]int{0, 0}] = 0

	var curX, curY int = 0, 0
	var minX, minY int = curX, curY
	var maxX, maxY int = curX, curY
	curDrcIdx := 0
	for i := 0; i < K; i++ {
		v, ok := board[[2]int{curX, curY}]
		if !ok || v == 0 { // 当前格子白色
			board[[2]int{curX, curY}] = (board[[2]int{curX, curY}] + 1) % 2 // 翻转颜色
			curDrcIdx = (curDrcIdx + 1) % 4                                 // 更改方向
		} else { // 当前格子黑色
			board[[2]int{curX, curY}] = (board[[2]int{curX, curY}] + 1) % 2 // 翻转颜色
			curDrcIdx = (curDrcIdx + 3) % 4                                 // 更改方向
		}

		// 走一步
		switch curDrcIdx {
		case 0:
			curY++
		case 1:
			curX++
		case 2:
			curY--
		case 3:
			curX--
		}

		if minX > curX {
			minX = curX
		}

		if minY > curY {
			minY = curY
		}

		if maxX < curX {
			maxX = curX
		}

		if maxY < curY {
			maxY = curY
		}
	}

	bd := make([][]byte, maxX-minX+1)
	for i := range bd {
		bd[i] = make([]byte, maxY-minY+1)
	}

	for i := 0; i < len(bd); i++ {
		for j := 0; j < len(bd[0]); j++ {
			if v, ok := board[[2]int{i + minX, j + minY}]; !ok {
				bd[i][j] = '_'
			} else {
				bd[i][j] = color[v]
			}
		}
	}

	bd[curX-minX][curY-minY] = direction[curDrcIdx]

	ret := make([]string, len(bd))
	for i := 0; i < len(bd); i++ {
		ret[i] = string(bd[i])
	}

	return ret
}
