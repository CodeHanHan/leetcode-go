package topic88

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	a := []int{1, 3, 0}
	merge(a, 2, []int{2}, 1)

	require.Equal(t, []int{1, 2, 3}, a)
}
