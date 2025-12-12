package utilities

import (
	"cmp"
)

func NumCombinations[T Integer](n, r T) uint64 {
	if cmp.Compare(r, n) == 1 {
		return 0
	}
	if cmp.Compare(r, n) == 0 {
		return 1
	}
	c := Factorial(n)
	den := Factorial(n - r)
	den *= Factorial(r)
	c /= den
	return c
}

func Combinations(n, m int, emit func([]int)) {
	s := make([]int, m)
	last := m - 1
	var rc func(int, int)
	rc = func(i, next int) {
		for j := next; j < n; j++ {
			s[i] = j
			if i == last {
				emit(s)
			} else {
				rc(i+1, j+1)
			}
		}
	}
	rc(0, 0)
}
