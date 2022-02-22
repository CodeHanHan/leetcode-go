package topic557

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_reverseWords(t *testing.T) {
	require.Equal(t, "s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))

	require.Equal(t, "aa bb cc", reverseWords("aa bb cc"))
}
