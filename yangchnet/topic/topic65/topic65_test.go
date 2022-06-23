package topic65

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isNumber(t *testing.T) {
	require.Equal(t, false, isNumber("e"))

	require.Equal(t, true, isNumber("1e7"))
}
