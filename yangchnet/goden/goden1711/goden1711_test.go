package goden1711

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findClosest(t *testing.T) {
	require.Equal(t, 1, findClosest(
		[]string{"I", "am", "a", "student", "from", "a", "university", "in", "a", "city"},
		"a", "student"))
}
