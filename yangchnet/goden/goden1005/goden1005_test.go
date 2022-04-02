package goden1005

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findString(t *testing.T) {
	require.Equal(t, -1, findString([]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ta"))
	require.Equal(t, 4, findString([]string{"at", "", "", "", "ball", "", "", "car", "", "", "dad", "", ""}, "ball"))
}
