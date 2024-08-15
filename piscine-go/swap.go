package piscine

func Swap(a *int, b *int) {
	k := *b
	z := *a
	*a =k
	*b = z

}