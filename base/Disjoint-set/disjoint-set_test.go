package disjointset

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRelative(t *testing.T) {
	testcases := []struct {
		n    int
		p    [][]int
		q    [][]int
		want []bool
	}{
		{
			n:    3,
			p:    [][]int{{1, 2}, {1, 3}},
			q:    [][]int{{2, 3}},
			want: []bool{true},
		},
		{
			n:    4,
			p:    [][]int{{1, 2}, {1, 3}},
			q:    [][]int{{2, 3}, {1, 4}, {2, 4}},
			want: []bool{true, false, false},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.want, findRelative(tt.n, tt.p, tt.q))
	}
}
