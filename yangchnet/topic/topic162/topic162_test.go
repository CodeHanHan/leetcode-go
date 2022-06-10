package topic162

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findPeakElement(t *testing.T) {
	require.Equal(t, 2, findPeakElement([]int{1, 2, 3, 1}))
}
