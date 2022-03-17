package goden0804

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_subsets(t *testing.T) {
	testcases := []struct {
		in  []int
		out [][]int
	}{
		{
			in:  []int{0},
			out: [][]int{{0}, {}},
		},
		{
			in:  []int{0, 1},
			out: [][]int{{0, 1}, {0}, {1}, {}},
		},
		{
			in:  []int{1, 2, 3},
			out: [][]int{{1, 2, 3}, {1, 2}, {1, 3}, {1}, {2, 3}, {2}, {3}, {}},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.out, subsets(tt.in))
	}
}
