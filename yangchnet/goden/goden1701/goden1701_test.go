package goden1701

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_add(t *testing.T) {
	require.Equal(t, 3, add(1, 2))
	require.Equal(t, 26, add(8, 18))
}
