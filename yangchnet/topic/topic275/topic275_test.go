package topic275

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_hIndex(t *testing.T) {
	require.Equal(t, 3, hIndex([]int{0, 1, 3, 5, 6}))
}
