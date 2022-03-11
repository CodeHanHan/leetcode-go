package dynamic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_FrogSteps(t *testing.T) {
	require.Equal(t, 2, FrogSteps(2))

	require.Equal(t, 21, FrogSteps(7))

	require.Equal(t, 1, FrogSteps(0))
}
