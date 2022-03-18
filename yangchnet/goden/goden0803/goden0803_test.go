package goden0803

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findMagicIndex(t *testing.T) {
	cases := []struct {
		in  []int
		out int
	}{
		{
			in:  []int{0, 1, 2},
			out: 0,
		},
		{
			in:  []int{1, 2, 3},
			out: -1,
		},
	}
	for _, c := range cases {
		require.Equal(t, c.out, findMagicIndex(c.in))
	}
}
