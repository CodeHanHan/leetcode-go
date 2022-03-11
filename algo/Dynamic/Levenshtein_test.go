package dynamic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Levenshtein(t *testing.T) {
	require.Equal(t, 3, LevenshteinDist("mtacnu", "mitcmu"))
}
