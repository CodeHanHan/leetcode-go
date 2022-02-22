package topic43

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_multiply(t *testing.T) {
	require.Equal(t, "5162", multiply("58", "89"))

	require.Equal(t, "40", multiply("5", "8"))

	require.Equal(t, "0", multiply("0", "8"))

	require.Equal(t, "89899011", multiply("999", "89989"))
}
