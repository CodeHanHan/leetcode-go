package goden1726

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_computeSimilarities(t *testing.T) {
	require.Equal(t,
		[]string{
			"0,1: 0.2500",
			"0,2: 0.1000",
			"2,3: 0.1429",
		},
		computeSimilarities(
			[][]int{
				{14, 15, 100, 9, 3},
				{32, 1, 9, 3, 5},
				{15, 29, 2, 6, 8, 7},
				{7, 10},
			},
		),
	)
}

// func Test_diffAndSame(t *testing.T) {
// 	var diff, same int

// 	diff, same = diffAndSame([]int{3, 9, 14, 15, 100}, []int{1, 3, 5, 9, 32})
// 	require.Equal(t, 6, diff)
// 	require.Equal(t, 2, same)

// 	diff, same = diffAndSame([]int{2, 6, 7, 8, 15, 29}, []int{7, 10})
// 	require.Equal(t, 6, diff)
// 	require.Equal(t, 1, same)
// }
