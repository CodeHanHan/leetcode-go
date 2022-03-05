package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Graph(t *testing.T) {
	g := NewGraphByMatrix([][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 1, 0, 0},
		{1, 0, 0, 1, 1},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 1, 0},
	})

	dfs := make([]int, 0)
	g.DFS(func(n int) {
		dfs = append(dfs, n)
	})
	require.Equal(t, []int{0, 1, 2, 3, 4}, dfs)

	bfs := make([]int, 0)
	g.BFS(func(n int) {
		bfs = append(bfs, n)
	})
	require.Equal(t, []int{0, 1, 2, 3, 4}, bfs)
}
