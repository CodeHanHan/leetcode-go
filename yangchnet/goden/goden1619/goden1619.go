package goden1619

import "sort"

func pondSizes(land [][]int) []int {
	res := make([]int, 0)
	for i := 0; i < len(land); i++ {
		for j := 0; j < len(land[0]); j++ {
			if land[i][j] == 0 {
				res = append(res, expand(i, j, land))
			}
		}
	}

	sort.Ints(res)

	return res
}

func expand(i, j int, land [][]int) int {
	// 越界
	if i < 0 || j < 0 || i >= len(land) || j >= len(land[0]) {
		return 0
	}

	// 已访问过
	if land[i][j] != 0 {
		return 0
	}

	land[i][j] = -1

	// 未访问过且为水域
	return 1 +
		expand(i+1, j, land) +
		expand(i-1, j, land) +
		expand(i, j+1, land) +
		expand(i, j-1, land) +
		expand(i+1, j+1, land) +
		expand(i+1, j-1, land) +
		expand(i-1, j+1, land) +
		expand(i-1, j-1, land)

}
