package piscine

func DescendAppendRange(min, max int) []int {
	gaga := []int{}
	if min < max {
		return gaga
	}

	for i := min; i > max; i-- {
		gaga = append(gaga, i)
	}
	return gaga
}
