package goden0807

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_permutation(t *testing.T) {
	testcases := []struct {
		in  string
		out []string
	}{
		{
			in:  "a",
			out: []string{"a"},
		},

		{
			in:  "ab",
			out: []string{"ab", "ba"},
		},
		{
			in:  "qwe",
			out: []string{"qwe", "qew", "wqe", "weq", "eqw", "ewq"},
		},
		{
			in:  "xFr",
			out: []string{"xFr", "xrF", "Fxr", "Frx", "rxF", "rFx"},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.out, permutation(tt.in))
	}
}
