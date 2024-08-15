package piscine

func Max(slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	max := slice[0]
	for _, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}
