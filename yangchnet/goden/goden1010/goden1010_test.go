package goden1010

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_StreamRank(t *testing.T) {
	sr := Constructor()
	require.Equal(t, 0, sr.GetRankOfNumber(1))
	sr.Track(0)
	require.Equal(t, 1, sr.GetRankOfNumber(1))
	sr.Track(1)
	require.Equal(t, 2, sr.GetRankOfNumber(1))
}
