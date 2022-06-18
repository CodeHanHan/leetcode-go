package topic50

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_myPow(t *testing.T) {
	require.Equal(t, 4.0, myPow(2, 2))
	require.Equal(t, 0.25, myPow(2, -2))
}
