package goden1720

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Constructor(t *testing.T) {
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	// obj.FindMedian()
	require.Equal(t, 1.5, obj.FindMedian())
	obj.AddNum(3)
	require.Equal(t, 2.0, obj.FindMedian())
	obj.AddNum(4)
	require.Equal(t, 2.5, obj.FindMedian())
}
