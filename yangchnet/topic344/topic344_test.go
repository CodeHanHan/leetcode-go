package topic344

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseString(t *testing.T) {
	a := []byte{1, 2}
	reverseString(a)
	require.Equal(t, []byte{2, 1}, a)

	b := []byte{1, 2, 3, 4, 5}
	reverseString(b)
	require.Equal(t, []byte{5, 4, 3, 2, 1}, b)
}
