package topic006

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_convert(t *testing.T) {
	require.Equal(t, "PAHNAPLSIIGYIR", convert("PAYPALISHIRING", 3))

	require.Equal(t, "PINALSIGYAHRPI", convert("PAYPALISHIRING", 4))
}

func Test_convert_1(t *testing.T) {
	require.Equal(t, "PAHNAPLSIIGYIR", convert_1("PAYPALISHIRING", 3))

	require.Equal(t, "PINALSIGYAHRPI", convert_1("PAYPALISHIRING", 4))
}
