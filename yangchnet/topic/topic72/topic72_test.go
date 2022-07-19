package topic72

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDistance(t *testing.T) {
	require.Equal(t, 3, minDistance("horse", "ros"))

	require.Equal(t, 2, minDistance("abb", "bba"))

	require.Equal(t, 5, minDistance("intention", "execution"))

	require.Equal(t, 3, minDistance("zoolog", "zoogeo"))

	require.Equal(t, 2, minDistance("ab", "abbb"))
}
