package lc217

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_containsDuplicate(t *testing.T) {

	require.Equal(t, true, containsDuplicate([]int{1, 2, 3, 1}))

	require.Equal(t, false, containsDuplicate([]int{1, 2, 3, 4}))

	require.Equal(t, true, containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}

func Test_containsDuplicate_1(t *testing.T) {

	require.Equal(t, true, containsDuplicate([]int{1, 2, 3, 1}))

	require.Equal(t, false, containsDuplicate([]int{1, 2, 3, 4}))

	require.Equal(t, true, containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
