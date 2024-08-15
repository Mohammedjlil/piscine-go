package piscine

func HashCode(dec string) string {
	do := ""
	for _, v := range dec {
		fo := int(v)
		if v >= 32 && v <= 127 {
			do += string(rune((fo + len(dec)) % 126))
		} else {
			do += string(rune(((fo + len(dec)) % 126) + 33))
		}
	}
	return do
}
