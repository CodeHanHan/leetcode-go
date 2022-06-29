package topic96

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numTrees(t *testing.T) {
	require.Equal(t, 5, numTrees(3))

	require.Equal(t, 1, numTrees(1))
}
