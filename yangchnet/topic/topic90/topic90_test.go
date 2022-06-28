package topic90

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsetsWithDup(t *testing.T) {
	require.ElementsMatch(t, [][]int{
		{},
		{1},
		{2},
		{3},
		{1, 2},
		{1, 3},
		{2, 3},
		{1, 2, 3},
	}, subsetsWithDup([]int{1, 2, 3}))

	require.ElementsMatch(t, [][]int{
		{},
		{1},
		{2},
		{1, 2},
		{2, 2},
		{1, 2, 2},
	}, subsetsWithDup([]int{1, 2, 2}))

	require.ElementsMatch(t, [][]int{
		{},
		{2},
		{2, 2},
	}, subsetsWithDup([]int{2, 2}))

	require.ElementsMatch(t, [][]int{
		{},
		{2},
	}, subsetsWithDup([]int{2}))
}
