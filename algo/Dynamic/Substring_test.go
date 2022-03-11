package dynamic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Substring(t *testing.T) {
	require.Equal(t, 4, LongestCommonSubString("flower", "flowa"))
	require.Equal(t, 4, LongestCommonSubString("aaflower", "flowa"))
}
