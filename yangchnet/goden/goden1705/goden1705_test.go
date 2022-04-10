package goden1705

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findLongestSubarray(t *testing.T) {
	require.Equal(t,
		[]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7"},
		findLongestSubarray(
			[]string{"A", "1", "B", "C", "D", "2", "3", "4", "E", "5", "F", "G", "6", "7", "H", "I", "J", "K", "L", "M"}))
}
