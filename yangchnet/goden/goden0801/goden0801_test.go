package goden0801

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_waysToStep(t *testing.T) {
	// require.Equal(t, 13, waysToStep(5))
	// require.Equal(t, 4, waysToStep(3))
	require.Equal(t, 1, waysToStep(1))
}
