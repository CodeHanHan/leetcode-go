package topic155

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Constructor(t *testing.T) {
	s := Constructor()

	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	require.Equal(t, -3, s.GetMin())
	s.Pop()
	require.Equal(t, 0, s.Top())
	require.Equal(t, -2, s.GetMin())
}

func Test_Constructor1(t *testing.T) {
	s := Constructor()

	s.Push(1)
	s.Push(2)
	require.Equal(t, 2, s.Top())
	require.Equal(t, 1, s.GetMin())
	s.Pop()
	require.Equal(t, 1, s.GetMin())
	require.Equal(t, 1, s.Top())
}
