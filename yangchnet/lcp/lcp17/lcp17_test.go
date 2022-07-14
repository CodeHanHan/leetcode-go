package lcp17

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_calculate(t *testing.T) {
	require.Equal(t, 4, calculate("AB"))
}
