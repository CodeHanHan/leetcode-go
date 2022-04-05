package goden1618

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_patternMatching(t *testing.T) {
	require.Equal(t, true, patternMatching("abba", "dogcatcatdog"))
	require.Equal(t, false, patternMatching("abba", "dogcatcatfish"))
	require.Equal(t, false, patternMatching("aaaa", "dogcatcatdog"))
	require.Equal(t, true, patternMatching("abba", "dogcatcatdog"))
	require.Equal(t, false, patternMatching("abbba", "dogdogdogdogdog"))
}

func Test_canDivide(t *testing.T) {
	require.Equal(t, true, canDivide("aaa", 3))
	require.Equal(t, false, canDivide("aaa", 2))
	require.Equal(t, true, canDivide("catcat", 2))
	require.Equal(t, false, canDivide("catcat", 3))
}
