package topic295

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Constructor(t *testing.T) {
	f := Constructor()

	f.AddNum(1)
	f.AddNum(2)
	f.AddNum(3)
	f.AddNum(4)

	require.Equal(t, 2.5, f.FindMedian())

	f.AddNum(5)
	require.Equal(t, 3.0, f.FindMedian())
}
