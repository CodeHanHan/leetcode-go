package topic60

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getPermutation(t *testing.T) {
	require.Equal(t, "213", getPermutation(3, 3))
}
