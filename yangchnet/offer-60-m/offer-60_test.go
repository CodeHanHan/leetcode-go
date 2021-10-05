package offer_60_m

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_dicesProbability(t *testing.T) {
	require.Equal(t,
		[]float64{0.16667, 0.16667, 0.16667, 0.16667, 0.16667, 0.16667},
		dicesProbability(1))

	require.Equal(t,
		[]float64{0.02778, 0.05556, 0.08333, 0.11111, 0.13889, 0.16667, 0.13889, 0.11111, 0.08333, 0.05556, 0.02778},
		dicesProbability(2))

	require.Equal(t,
		[]float64{0.16667, 0.16667, 0.16667, 0.16667, 0.16667, 0.16667},
		dicesProbability1(1))

	require.Equal(t,
		[]float64{0.02778, 0.05556, 0.08333, 0.11111, 0.13889, 0.16667, 0.13889, 0.11111, 0.08333, 0.05556, 0.02778},
		dicesProbability1(2))
}
