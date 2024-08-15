package piscine

func Abort(a, b, c, d, e int) int {
	namber := [5]int{a, b, c, d, e}
	for i := 0; i < 5; i++ {
		for j := i + 1; j < 5; j++ {
			if namber[i] > namber[j] {
				namber[i], namber[j] = namber[j], namber[i]
			}
		}
	}
	return namber[2]
}
