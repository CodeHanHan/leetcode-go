package topic030

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findSubstring(t *testing.T) {
	require.Equal(t, []int{0, 9}, findSubstring("barfoothefoobarman", []string{"foo", "bar"}))

	require.Equal(t, []int{}, findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))

	require.Equal(t, []int{6, 9, 12}, findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))

	require.Equal(t, []int{8}, findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))

	require.Equal(t, []int{0, 3, 6}, findSubstring("foobarfoobar", []string{"foo", "bar"}))
}

func Test_findSubstring_1(t *testing.T) {
	require.Equal(t, []int{0, 9}, findSubstring_1("barfoothefoobarman", []string{"foo", "bar"}))

	require.Equal(t, []int{}, findSubstring_1("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))

	require.Equal(t, []int{6, 9, 12}, findSubstring_1("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))

	require.Equal(t, []int{8}, findSubstring_1("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))

	require.Equal(t, []int{0, 3, 6}, findSubstring_1("foobarfoobar", []string{"foo", "bar"}))
}
