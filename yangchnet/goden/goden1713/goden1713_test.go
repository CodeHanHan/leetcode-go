package goden1713

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_respace(t *testing.T) {
	require.Equal(t, 7, respace(
		[]string{"looked", "just", "like", "her", "brother"},
		"jesslookedjustliketimherbrother"))

	require.Equal(t, 5, respace(
		[]string{"bt", "vbtbvtvvbbvtbvvbbbvbtbbv", "bvvbbbvvvbvttbtbvtvtvvbttbbbtvvvb", "btttbvbbbtbbtbvvttbvbvtvbtbbttb", "bt", "vvbvbvbvtbvbvvvvtv", "tbvvvtttvtbvbtttvvbtvvvvtvvbvvtvvbbvbbbvb", "v", "bvb", "vvtbvtvbttbttvvbvttbt", "btbtbtttvbbtbttbtvvttbvtbtvtbvvtbbbb", "bttbvbbttvvbtvvvvb", "bvvbvbvttbvtbvvtbttvvvvtvbtvbttbbvvtvtvv", "vbtttt", "btbvbbbvbtvtbvvv", "b", "tbtbtv", "vbvbbvvbvbbvvb", "vbvvtvbvbvbttvbvbtvbtbtvtbvbbtb", "bvvbvvttttttbtvvvttvbvtvvbvtbtvtbvttvvtbt", "bvtbttv", "bbbt", "vtt", "ttvv", "b", "vbb", "vtvvbtttvtbbvvbbtbb", "vvv", "vvvvbbbtbbbvbbbb", "ttvvbtvv", "v", "btvbbvtbbvbvtvttvvbbbtbvvvtbtb", "vv", "btbttbtbbvbvvbvttbttvtbbtbvtttvbbtbbtvtvvvvbbttvtvvbbbv", "ttvbbbbttbtbtb", "tvvbvbvvb", "vv", "ttbttvtvbtbbbbv", "bvvvtbbvvvbtvbvtvtvvvvbb", "vtbvvbvvvvvttvbbttbbvtvt", "tbvbbt", "b", "v", "tvbbtvvtvvtbtbttvvvb", "tbttbttbbbtbtvtvtvtbbbvvtbbbvbbvvvbbttvvt", "bbvttvtvvtbvbbttvbbtbvvttbvbvbtbvvvbbbvbvbvbtvbtvvvtvvtbttbttbbvbbbttvvvbvvtb", "vttvvbvv", "tbttbvvttvbtvvtbbvvv", "bvbbbbbbbb", "vtbvvtbbvtt", "bvttbvvbvb", "tvttttbbvvvbbtttvvv"},
		"btbvtttttbvttbvvbbtvvbvbvvbtbtbtvbtttbvbbbtbbtbvvttbvbvtvbtbbttbvvbvbtttbvttbvvbbvvv"))

	require.Equal(t, 7, respace(
		[]string{"looked", "just", "like", "her", "brother", "a"},
		"jessalookedjustliketimherbrother"))

	// require.Equal(t, 7, respace(
	// 	[]string{"aaysaayayaasyya", "yyas", "yayysaaayasasssy", "yaasassssssayaassyaayaayaasssasysssaaayysaaasaysyaasaaaaaasayaayayysasaaaa", "aya", "sya", "ysasasy", "syaaaa", "aaaas", "ysa", "a", "aasyaaassyaayaayaasyayaa", "ssaayayyssyaayyysyayaasaaa", "aya", "aaasaay", "aaaa", "ayyyayssaasasysaasaaayassasysaaayaassyysyaysaayyasayaaysyyaasasasaayyasasyaaaasysasy", "aaasa", "ysayssyasyyaaasyaaaayaaaaaaaaassaaa", "aasayaaaayssayyaayaaaaayaaays", "s"},
	// 	"asasayaayaassayyayyyyssyaassasaysaaysaayaaaaysyaaaa",
	// ))
}
