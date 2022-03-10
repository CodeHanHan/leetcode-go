package goden0502

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_printBin(t *testing.T) {
	require.Equal(t, "0.101", printBin(0.625))
	require.Equal(t, "ERROR", printBin(0.1))
}
