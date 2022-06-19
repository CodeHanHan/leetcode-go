package topic57

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_insert(t *testing.T) {
	require.Equal(t, [][]int{{1, 5}, {6, 9}}, insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}))
}
