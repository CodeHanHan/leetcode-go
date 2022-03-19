package goden0808

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
			out: []string{"eqw", "ewq", "qew", "qwe", "weq", "wqe"},
		},
		{
			in:  "xFr",
			out: []string{"Frx", "Fxr", "rFx", "rxF", "xFr", "xrF"},
		},
		{
			in:  "aa",
			out: []string{"aa"},
		},
		{
			in:  "qqe",
			out: []string{"eqq", "qeq", "qqe"},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.out, permutation(tt.in))
	}
}
