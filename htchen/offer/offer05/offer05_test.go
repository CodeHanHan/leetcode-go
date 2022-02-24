package offer05

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_replaceSpace(t *testing.T) {
	require.Equal(t, "We%20are%20happy.", replaceSpace("We are happy."))

	require.Equal(t, "i%20love%20%20you.", replaceSpace("i love  you."))
}
