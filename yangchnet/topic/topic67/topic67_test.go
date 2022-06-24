package topic67

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addBinary(t *testing.T) {
	require.Equal(t, "11", addBinary("00", "11"))

	require.Equal(t, "111", addBinary("110", "1"))

	require.Equal(t, "001", addBinary("000", "1"))

	require.Equal(t, "1000", addBinary("101", "11"))

	require.Equal(t, "10101", addBinary("1010", "1011"))

	require.Equal(t, "110001", addBinary("101111", "10"))
}
