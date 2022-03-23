package goden0811

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_waysToChange(t *testing.T) {
	// require.Equal(t, 2, waysToChange(5))

	require.Equal(t, 4, waysToChange(10))
}
