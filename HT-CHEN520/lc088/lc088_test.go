package lc088

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_merge(t *testing.T) {
	prea, preb := []int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}
	merge(prea, 3, preb, 3)
	require.Equal(t, []int{1, 2, 2, 3, 5, 6}, prea)

	prec, pred := []int{1}, []int{}
	merge(prec, 1, pred, 0)
	require.Equal(t, []int{1}, prec)
}
