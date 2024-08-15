package piscine

func ShoppingListSort(list []string) []string {
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if len(list[i]) > len(list[j]) {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}
