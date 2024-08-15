package piscine

func Unmatch(a []int) int {
	gool := 0

	for _, fo := range a {
		for _, salam := range a {
			if fo == salam {
				gool++
			}
		}
		if gool%2 != 0 {
			return fo
		}
	}

	return -1
}
