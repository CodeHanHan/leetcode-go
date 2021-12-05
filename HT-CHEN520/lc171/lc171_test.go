package lc171

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_titleToNumber(t *testing.T) {

	require.Equal(t, 1, titleToNumber("A"))

	require.Equal(t, 28, titleToNumber("AB"))

	require.Equal(t, 701, titleToNumber("ZY"))

	require.Equal(t, 2147483647, titleToNumber("FXSHRXW"))
}
