package topic115

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_numDistinct(t *testing.T) {
	require.Equal(t, 3, numDistinct("rabbbit", "rabbit"))
}
