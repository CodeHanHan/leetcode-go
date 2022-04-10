package goden1706

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numberOf2sInRange(t *testing.T) {
	require.Equal(t, 1, numberOf2sInRange(2))
	require.Equal(t, 2, numberOf2sInRange(12))
	require.Equal(t, 2, numberOf2sInRange(13))
}

func Test_numberOf2sInRange2(t *testing.T) {
	require.Equal(t, 1, numberOf2sInRange2(2))
	require.Equal(t, 2, numberOf2sInRange2(12))
	require.Equal(t, 2, numberOf2sInRange2(13))
	require.Equal(t, 9, numberOf2sInRange2(25))
}
