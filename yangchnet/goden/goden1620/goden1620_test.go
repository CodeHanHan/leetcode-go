package goden1620

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_getValidT9Words(t *testing.T) {
	require.Equal(t, []string{"tree", "used"}, getValidT9Words("8733", []string{"tree", "used"}))
	require.Equal(t, []string{"a", "b", "c"}, getValidT9Words("2", []string{"a", "b", "c"}))
}

func Test_getValidT9Words2(t *testing.T) {
	require.Equal(t, []string{"tree", "used"}, getValidT9Words2("8733", []string{"tree", "used"}))
	require.Equal(t, []string{"a", "b", "c"}, getValidT9Words2("2", []string{"a", "b", "c"}))
}
