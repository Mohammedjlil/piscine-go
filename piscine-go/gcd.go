package piscine

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0

	}
	for b != 0 {
		x := a % b 
		a , b = b , x 
	}
	return a

}