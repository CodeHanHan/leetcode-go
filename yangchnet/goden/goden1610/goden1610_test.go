package goden1610

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_maxAliveYear(t *testing.T) {
	require.Equal(t, 1901, maxAliveYear([]int{1900, 1901, 1950}, []int{1948, 1951, 2000}))
	require.Equal(t, 1960, maxAliveYear(
		[]int{1972, 1908, 1915, 1957, 1960, 1948, 1912, 1903, 1949, 1977, 1900, 1957, 1934, 1929, 1913, 1902, 1903, 1901},
		[]int{1997, 1932, 1963, 1997, 1983, 2000, 1926, 1962, 1955, 1997, 1998, 1989, 1992, 1975, 1940, 1903, 1983, 1969}))
}

func Test_maxAliveYear2(t *testing.T) {
	require.Equal(t, 1901, maxAliveYear2([]int{1900, 1901, 1950}, []int{1948, 1951, 2000}))
	require.Equal(t, 1960, maxAliveYear2(
		[]int{1972, 1908, 1915, 1957, 1960, 1948, 1912, 1903, 1949, 1977, 1900, 1957, 1934, 1929, 1913, 1902, 1903, 1901},
		[]int{1997, 1932, 1963, 1997, 1983, 2000, 1926, 1962, 1955, 1997, 1998, 1989, 1992, 1975, 1940, 1903, 1983, 1969}))
}
