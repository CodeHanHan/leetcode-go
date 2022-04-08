package goden1626

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculate(t *testing.T) {
	require.Equal(t, 2, calculate("1+ 1"))
	require.Equal(t, 3, calculate("1+ 1*2"))
	require.Equal(t, 12, calculate("2*5+ 1*2"))
	require.Equal(t, 20, calculate("2*5/1*2"))
	require.Equal(t, 6, calculate("2*5/3*2"))
	require.Equal(t, 1, calculate(" 3/2"))
	require.Equal(t, 32, calculate(" 32"))
	require.Equal(t, -24, calculate("1*2-3/4+5*6-7*8+9/10"))
	require.Equal(t, -24, calculate("2-0+30-56+0"))
}
