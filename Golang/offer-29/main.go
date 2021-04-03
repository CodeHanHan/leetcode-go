package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	flag := 0
	walk := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	visited := make([][]bool, len(matrix))
	for i := 0; i < len(visited); i++ {
		visited[i] = make([]bool, len(matrix[0]))
	}

	cur := []int{0, 0}
	result := []int{}
	//result = append(result, matrix[0][0])
	for {
		visited[cur[0]][cur[1]] = true
		result = append(result, matrix[cur[0]][cur[1]])
		if cross(matrix, cur, walk[flag]) || isvisited(visited, cur, walk[flag]) {
			flag = (flag + 1) % 4
		}
		if isvisited(visited, cur, walk[flag]) || cross(matrix, cur, walk[flag]) {
			return result
		}

		cur[0] += walk[flag][0]
		cur[1] += walk[flag][1]
	}
}

func cross(matrix [][]int, cur []int, walk []int) bool {
	cur_cow := cur[0] + walk[0]
	cur_column := cur[1] + walk[1]
	if cur_cow < 0 || cur_cow > len(matrix)-1 {
		return true
	}
	if cur_column < 0 || cur_column > len(matrix[0])-1 {
		return true
	}
	return false
}

func isvisited(visited [][]bool, cur []int, walk []int) bool {
	cur_cow := cur[0] + walk[0]
	cur_column := cur[1] + walk[1]
	if cur_cow < 0 || cur_cow > len(visited)-1 {
		return false
	}
	if cur_column < 0 || cur_column > len(visited[0])-1 {
		return false
	}
	return visited[cur_cow][cur_column]
}
