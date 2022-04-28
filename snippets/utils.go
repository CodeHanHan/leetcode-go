package snippets

import (
	"golang.org/x/exp/constraints"
)

func MinN[T constraints.Ordered](vars ...T) T {
	return most(min[T], vars...)
}

func MaxN[T constraints.Ordered](vars ...T) T {
	return most(max[T], vars...)
}

func most[T constraints.Ordered](cmp func(a, b T) T, vars ...T) T {
	var m T
	if len(vars) <= 0 {
		return m
	}

	m = vars[0]
	for i := 1; i < len(vars); i++ {
		m = cmp(m, vars[i])
	}

	return m
}

func min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func Abs[T interface {
	constraints.Signed | constraints.Float
}](x T) T {
	if x < 0 {
		return -x
	}

	return x
}
