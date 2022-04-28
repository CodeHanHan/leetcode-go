package snippets

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minN(t *testing.T) {
	require.Equal(t, 1, MinN(1, 2))
	require.Equal(t, "a", MinN("a", "b"))
	require.Equal(t, 1.1, MinN(1.1, 1.2))
}
