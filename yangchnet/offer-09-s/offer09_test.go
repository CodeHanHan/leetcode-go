package offer_09_s

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	q := Constructor()

	q.AppendTail(1)
	q.AppendTail(2)
	q.AppendTail(3)

	require.NotNil(t, q)

	require.EqualValues(t, q.DeleteHead(), 1)
	require.EqualValues(t, q.DeleteHead(), 2)
	require.EqualValues(t, q.DeleteHead(), 3)
}
