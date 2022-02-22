package topic89

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_grayCode(t *testing.T) {
	require.Equal(t, []int{0, 1, 3, 2, 6, 7, 5, 4}, grayCode(3))

	require.Equal(t, []int{0, 1, 3, 2}, grayCode(2))

	require.Equal(t, []int{0}, grayCode(0))
}
