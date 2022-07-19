package topic232

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Constructor(t *testing.T) {
	q := Constructor()

	require.Equal(t, true, q.Empty())
	q.Push(1)
	q.Push(2)
	require.Equal(t, 1, q.Peek())

	require.Equal(t, 1, q.Pop())

	require.Equal(t, false, q.Empty())
}
