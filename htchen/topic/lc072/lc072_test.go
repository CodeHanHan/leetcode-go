package lc072

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minDistance(t *testing.T) {
	require.Equal(t, 3, minDistance("horse", "ros"))
	
    require.Equal(t, 5, minDistance("intention", "execution"))
}
