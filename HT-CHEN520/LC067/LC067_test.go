package lc067

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addBinary(t *testing.T) {

	require.Equal(t, "100", addBinary("11", "1"))

	require.Equal(t, "10101", addBinary("1010", "1011"))
}
