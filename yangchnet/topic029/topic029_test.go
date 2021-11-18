package topic029

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_divide(t *testing.T) {
	require.Equal(t, 3, divide(10, 3))

	require.Equal(t, -2, divide(7, -3))
}
