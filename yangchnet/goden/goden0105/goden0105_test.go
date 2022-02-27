package goden0105

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_oneEditAway(t *testing.T) {
	require.Equal(t, true, oneEditAway("pale", "ple"))
	require.Equal(t, false, oneEditAway("pales", "ple"))
	require.Equal(t, true, oneEditAway("", "p"))
	require.Equal(t, true, oneEditAway("p", "p"))
	require.Equal(t, true, oneEditAway("abac", "abad"))

}
