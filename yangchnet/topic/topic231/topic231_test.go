package topic231

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPowerOfTwo(t *testing.T) {
	require.Equal(t, true, isPowerOfTwo(1))

	require.Equal(t, true, isPowerOfTwo(16))

	require.Equal(t, false, isPowerOfTwo(3))

	require.Equal(t, true, isPowerOfTwo(4))

	require.Equal(t, false, isPowerOfTwo(-1))

	require.Equal(t, false, isPowerOfTwo(0))

	require.Equal(t, true, isPowerOfTwo(int(math.Pow(float64(2), float64(31)))))

	require.Equal(t, false, isPowerOfTwo(-int(math.Pow(float64(2), float64(31)))))
}
