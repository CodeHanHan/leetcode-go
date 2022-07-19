package topiclc415

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_addStrings(t *testing.T) {
	require.Equal(t, "134", addStrings("11", "123"))

	require.Equal(t, "533", addStrings("456", "77"))
	
    require.Equal(t, "0", addStrings("0", "0"))
}
