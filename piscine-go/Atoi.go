package piscine

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	} 
	bo := 1
	to := 0

	if s[0] == '-' {
		bo = -1
		to =1
	} else if s[0] == '+' {
		to = 1
	}
	resolt := 0 
	for i := to ; i < len(s) ; i++ {
		if s[i] < '0' || '9' > s[i] {
			return 0
		}
		resolt = resolt * 10 + int(s[i] + '0') 
	}
	return bo * resolt
}