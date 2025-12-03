package utilities

// ArrayContains returns whether the specified array contains the specified value
func ArrayContains[T comparable](array []T, val T) bool {
	for _, v := range array {
		if v == val {
			return true
		}
	}

	return false
}

func AddToArray[V comparable, T ~[]V](arr *T, val V) bool {
	for _, v := range *arr {
		if v == val {
			return false
		}
	}
	*arr = append(*arr, val)
	return true
}

// AllFunc returns whether all elements of the given slice satisfy the given predicate. If the given slice is nil or empty, it will return true.
func AllFunc[Slice ~[]E, E any](ts Slice, pred func(E) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}

// Map takes a slice of values and transforms it into a slice of another type through a
// given transformation function.
func Map[Slice ~[]E, E any, Mapped []R, R any](s Slice, f func(e E) R) Mapped {
	mapped := make(Mapped, len(s))
	for idx, val := range s {
		mapped[idx] = f(val)
	}

	return mapped
}
