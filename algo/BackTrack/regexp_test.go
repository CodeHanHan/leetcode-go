package backtrack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Regexp(t *testing.T) {
	re := &Regexp{
		Pattern: "abd*",
	}

	require.Equal(t, true, re.Match("abdjfkajfa"))
	require.Equal(t, true, re.Match("abd"))

	re = &Regexp{
		Pattern: "ab?",
	}

	require.Equal(t, true, re.Match("abd"))
	require.Equal(t, true, re.Match("ab"))
	require.Equal(t, false, re.Match("abcd"))
}
