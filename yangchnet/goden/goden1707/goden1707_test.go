package goden1707

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_trulyMostPopular(t *testing.T) {
	require.ElementsMatch(t,
		[]string{"John(27)", "Chris(36)"},
		trulyMostPopular(
			[]string{"John(15)", "Jon(12)", "Chris(13)", "Kris(4)", "Christopher(19)"},
			[]string{"(Jon,John)", "(John,Johnny)", "(Chris,Kris)", "(Chris,Christopher)"},
		),
	)
}
