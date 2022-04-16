package lc038

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countAndSay(t *testing.T) {
	require.Equal(t, "1", countAndSay(1))
	
    require.Equal(t, "1211", countAndSay(4))
}
