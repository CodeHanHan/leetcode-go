package lc461

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hammingDistance1(t *testing.T) {
	require.Equal(t, 2, hammingDistance1(1, 4))

	require.Equal(t, 1, hammingDistance1(3, 1))
}

func Test_hammingDistance2(t *testing.T) {
	require.Equal(t, 2, hammingDistance2(1, 4))

	require.Equal(t, 1, hammingDistance2(3, 1))
}
