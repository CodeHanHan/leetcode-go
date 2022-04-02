package goden1003

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_search(t *testing.T) {
	require.Equal(t, 8, search2([]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 5))
	require.Equal(t, -1, search2([]int{15, 16, 19, 20, 25, 1, 3, 4, 5, 7, 10, 14}, 11))
}
