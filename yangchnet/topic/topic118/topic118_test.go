package topic118

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generate(t *testing.T) {
	require.Equal(t, [][]int{
		{1},
		{1, 1},
		{1, 2, 1},
		{1, 3, 3, 1},
		{1, 4, 6, 4, 1},
	}, generate(5))
}
