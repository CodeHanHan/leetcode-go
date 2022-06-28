package topic93

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_restoreIpAddresses(t *testing.T) {
	require.ElementsMatch(t, []string{"255.255.11.135", "255.255.111.35"}, restoreIpAddresses("25525511135"))

	require.ElementsMatch(t, []string{"0.0.0.0"}, restoreIpAddresses("0000"))

	require.ElementsMatch(t, []string{"0.1.2.201", "0.1.220.1", "0.12.20.1", "0.122.0.1"}, restoreIpAddresses("012201"))

	require.ElementsMatch(t, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}, restoreIpAddresses("101023"))

}

func Test_restoreIpAddresses1(t *testing.T) {
	require.ElementsMatch(t, []string{"255.255.11.135", "255.255.111.35"}, restoreIpAddresses1("25525511135"))

	require.ElementsMatch(t, []string{"0.0.0.0"}, restoreIpAddresses1("0000"))

	require.ElementsMatch(t, []string{"0.1.2.201", "0.1.220.1", "0.12.20.1", "0.122.0.1"}, restoreIpAddresses1("012201"))

	require.ElementsMatch(t, []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}, restoreIpAddresses1("101023"))
}
