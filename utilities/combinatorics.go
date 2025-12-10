package utilities

import (
	"cmp"
)

func Factorial[T Number](n T) T {
	if Sign(n) < 1 {
		return T(0)
	}
	r := T(1)
	i := T(2)
	for cmp.Compare(i, n) < 1 {
		r *= i
		i += T(1)
	}
	return r
}

func NumPermutations[T Number](n, k T) T {
	r := Factorial(n)
	r /= Factorial(n - k)
	return r
}

func NumCombinations[T Number](n, r T) T {
	if cmp.Compare(r, n) == 1 {
		return T(0)
	}
	if cmp.Compare(r, n) == 0 {
		return T(1)
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
