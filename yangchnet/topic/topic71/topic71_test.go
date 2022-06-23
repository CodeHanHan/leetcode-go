package topic71

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_simplifyPath(t *testing.T) {
	require.Equal(t, "/path", simplifyPath("path"))

	require.Equal(t, "/home", simplifyPath("/home/"))

	require.Equal(t, "/", simplifyPath("/../"))

	require.Equal(t, "/c", simplifyPath("/a/./b/../../c/"))
}
