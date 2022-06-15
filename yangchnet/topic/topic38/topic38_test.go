package topic38

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countAndSay(t *testing.T) {
	require.Equal(t, "1", countAndSay(1))
	require.Equal(t, "11", countAndSay(2))
	require.Equal(t, "21", countAndSay(3))
	require.Equal(t, "1211", countAndSay(4))
	require.Equal(t, "111221", countAndSay(5))
	require.Equal(t, "312211", countAndSay(6))
}
