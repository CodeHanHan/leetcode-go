package topic1143

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestCommonSubsequence(t *testing.T) {
	require.Equal(t, 3, longestCommonSubsequence("abcde", "ace"))
}
