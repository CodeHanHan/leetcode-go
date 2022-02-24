package LC014

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestCommonPrefix(t *testing.T) {

	require.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))

	require.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
