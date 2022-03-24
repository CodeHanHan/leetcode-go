package goden0810

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_floodFill(t *testing.T) {
	type param struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
	}

	testcases := []struct {
		in   param
		want [][]int
	}{
		{
			in: param{
				image: [][]int{
					{1, 1, 1},
					{1, 1, 0},
					{1, 0, 1},
				},
				sr:       1,
				sc:       1,
				newColor: 2,
			},
			want: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
		{
			in: param{
				image: [][]int{
					{1, 1, 1},
					{1, 1, 0},
					{1, 0, 1},
				},
				sr:       2,
				sc:       1,
				newColor: 2,
			},
			want: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 2, 1},
			},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.want, floodFill(tt.in.image, tt.in.sr, tt.in.sc, tt.in.newColor))
	}
}
