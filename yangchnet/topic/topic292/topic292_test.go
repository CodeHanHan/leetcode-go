package topic292

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canWinNim(t *testing.T) {
	require.Equal(t, false, canWinNim(4))

	require.Equal(t, true, canWinNim(1))

	require.Equal(t, true, canWinNim(2))

	require.Equal(t, false, canWinNim(100))

	require.Equal(t, true, canWinNim(131414111))
}
