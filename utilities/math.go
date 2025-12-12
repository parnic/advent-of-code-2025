package utilities

import "cmp"

func GCD[T Integer](a, b T) T {
	if b == 0 {
		return a
	}
	return GCD(b, a%b)
}

func LCM[T Integer](nums ...T) uint64 {
	num := len(nums)
	switch num {
	case 0:
		return 0
	case 1:
		return uint64(nums[0])
	}

	ret := lcm(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		ret = lcm(uint64(nums[i]), ret)
	}
	return ret
}

func lcm[T Integer](a, b T) uint64 {
	return uint64(a*b) / uint64(GCD(a, b))
}

func Sign[T Number](num T) int {
	if num == 0 {
		return 0
	} else if num > 0 {
		return 1
	}

	return -1
}

func Factorial[T Integer](n T) uint64 {
	if Sign(n) < 1 {
		return 0
	}
	r := uint64(1)
	i := uint64(2)
	for cmp.Compare(i, uint64(n)) < 1 {
		r *= i
		i += 1
	}
	return r
}
