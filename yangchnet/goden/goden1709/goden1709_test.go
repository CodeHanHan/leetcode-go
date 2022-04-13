package goden1709

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getKthMagicNumber(t *testing.T) {
	require.Equal(t, 9, getKthMagicNumber(5))
	require.Equal(t, 1, getKthMagicNumber(1))
}
