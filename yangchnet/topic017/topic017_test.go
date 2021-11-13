package topic017

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_letterCombinations(t *testing.T) {
	require.Equal(t, []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}, letterCombinations("23"))
}
