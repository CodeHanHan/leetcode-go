package goden1719

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_missingTwo(t *testing.T) {
	require.Equal(t, []int{1, 4}, missingTwo([]int{2, 3}))
}
