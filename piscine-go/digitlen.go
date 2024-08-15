package piscine

func DigitLen(n, base int) int {
	if n < 0 {
		n = -n
	}
	if base < 2 || base > 36 {
		return -1
	}

	c := 0
	for n > 0 {
		n = n / base
		c++

	}
	return c
}
