package topic165

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_compareVersion(t *testing.T) {
	require.Equal(t, 0, compareVersion("1.01", "1.001"))

	require.Equal(t, 0, compareVersion("1.0", "1.0.0"))

	require.Equal(t, -1, compareVersion("0.1", "1.1"))
}
