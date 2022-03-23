package goden0810

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] == newColor {
		return image
	}

	c := image[sr][sc]

	// 填充当前位置
	image[sr][sc] = newColor

	// 向上
	if sr-1 >= 0 && c == image[sr-1][sc] {
		floodFill(image, sr-1, sc, newColor)
	}

	// 向下
	if sr+1 < len(image) && c == image[sr+1][sc] {
		floodFill(image, sr+1, sc, newColor)
	}

	// 向左
	if sc-1 >= 0 && c == image[sr][sc-1] {
		floodFill(image, sr, sc-1, newColor)
	}

	// 向右
	if sc+1 < len(image[0]) && c == image[sr][sc+1] {
		floodFill(image, sr, sc+1, newColor)
	}

	return image
}
