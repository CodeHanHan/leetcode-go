package LC007

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverse(t *testing.T) {
	require.Equal(t, 321, reverse(123))

	require.Equal(t, -321, reverse(-123))

	require.Equal(t, 21, reverse(120))

	require.Equal(t, 0, reverse(0))

}
