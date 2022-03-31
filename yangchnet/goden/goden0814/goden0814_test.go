package goden0814

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_countEval(t *testing.T) {
	require.Equal(t, 2, countEval("1^0|0|1", 0))
}
