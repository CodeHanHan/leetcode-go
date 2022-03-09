package goden0501

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_insertBits(t *testing.T) {
	require.Equal(t, 1100, insertBits(1024, 19, 2, 6))
}
