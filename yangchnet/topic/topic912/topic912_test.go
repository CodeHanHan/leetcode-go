package topic912

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_sortArray(t *testing.T) {
	require.Equal(t, []int{1, 2, 3, 4, 5}, sortArray([]int{4, 1, 3, 5, 2}))
}
