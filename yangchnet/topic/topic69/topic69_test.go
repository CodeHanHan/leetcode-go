package topic69

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mySqrt(t *testing.T) {
	require.Equal(t, 0, mySqrt(0))
	require.Equal(t, 2, mySqrt(4))
	require.Equal(t, 1, mySqrt(1))
	require.Equal(t, 2, mySqrt(8))
	require.Equal(t, 3, mySqrt(14))
}
