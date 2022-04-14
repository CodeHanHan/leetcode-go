package lc135

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_candy(t *testing.T) {
	require.Equal(t, 5, candy([]int{1, 0, 2}))
	
    require.Equal(t, 4, candy([]int{1, 2, 2}))
}
