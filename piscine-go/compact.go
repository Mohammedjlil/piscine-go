package piscine

func Compact(ptr *[]string) int {
	coco := 0
	var slice []string

	for _, k := range *ptr {
		if k != "" {
			slice = append(slice, k)
			coco = coco + 1
		}
	}
	*ptr = slice
	return coco
}
