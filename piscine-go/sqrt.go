package piscine

func Sqrt(nb int) int {
	sqrt := 0
	for i := 1; i <= nb; i++ {
		sqrt = nb / i
		if sqrt == i && nb%i == 0 {
			return sqrt
		}
	}
	return 0
}
