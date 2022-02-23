package lc069

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mySqrt(t *testing.T) {

	require.Equal(t, 2, mySqrt(4))

	require.Equal(t, 2, mySqrt(8))

	require.Equal(t, 3, mySqrt(15))

	require.Equal(t, 1, mySqrt(1))
}
