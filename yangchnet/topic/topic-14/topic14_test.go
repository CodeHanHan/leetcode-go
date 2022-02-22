package topic11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_longestCommonPrefix(t *testing.T) {
	require.Equal(t, "fl", longestCommonPrefix([]string{"flower", "flow", "flight"}))

	require.Equal(t, "", longestCommonPrefix([]string{"dog", "racecar", "car"}))

	require.Equal(t, "", longestCommonPrefix([]string{}))

	require.Equal(t, "ajfjakfjajfjafakjf", longestCommonPrefix([]string{"ajfjakfjajfjafakjf"}))

	require.Equal(t, "ajfjakfjajfjafakjf", longestCommonPrefix([]string{"ajfjakfjajfjafakjf", "ajfjakfjajfjafakjf"}))

	require.Equal(t, "d", longestCommonPrefix([]string{"d", "dog"}))

	require.Equal(t, "d", longestCommonPrefix([]string{"dog", "d"}))
}

func Test_commonPrefix(t *testing.T) {
	require.Equal(t, "", commonPrefix("", ""))

	require.Equal(t, "", commonPrefix("a", "v"))

	require.Equal(t, "d", commonPrefix("dog", "d"))

	require.Equal(t, "d", commonPrefix("d", "dog"))

	require.Equal(t, "flow", commonPrefix("flower", "flow"))
}

func Test_longestCommonPrefix_1(t *testing.T) {
	require.Equal(t, "fl", longestCommonPrefix_1([]string{"flower", "flow", "flight"}))

	require.Equal(t, "", longestCommonPrefix_1([]string{"dog", "racecar", "car"}))

	require.Equal(t, "", longestCommonPrefix([]string{}))

	require.Equal(t, "ajfjakfjajfjafakjf", longestCommonPrefix_1([]string{"ajfjakfjajfjafakjf"}))

	require.Equal(t, "ajfjakfjajfjafakjf", longestCommonPrefix_1([]string{"ajfjakfjajfjafakjf", "ajfjakfjajfjafakjf"}))

	require.Equal(t, "d", longestCommonPrefix_1([]string{"d", "dog"}))

	require.Equal(t, "d", longestCommonPrefix_1([]string{"dog", "d"}))
}
