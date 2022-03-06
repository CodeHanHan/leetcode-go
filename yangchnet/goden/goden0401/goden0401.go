package goden0401

func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	if start == target {
		return true
	}

	for _, g := range graph {
		if g[1] == target {
			return findWhetherExistsPath(n, graph, start, g[0])
		}
	}

	return false
}
