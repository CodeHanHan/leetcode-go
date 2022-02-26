package goden0103

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_replaceSpaces(t *testing.T) {
	require.Equal(t, "Mr%20John%20Smith", replaceSpaces("Mr John Smith    ", 13))
	require.Equal(t, "%20%20%20%20%20", replaceSpaces("               ", 5))
}
