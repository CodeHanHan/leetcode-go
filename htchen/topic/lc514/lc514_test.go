package lc514

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRotateSteps(t *testing.T) {
	require.Equal(t, 4, findRotateSteps("godding", "gd"))

	require.Equal(t, 13, findRotateSteps("godding", "godding"))
}
