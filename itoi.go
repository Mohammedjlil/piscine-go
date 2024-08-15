package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}
	fi := ""
	if n < 0 {
		fi += "-"
		n =-n
	}
	to := n
	resolt := ""
	for n > 0 {
		n = n% 10
		resolt = string(n + '0') + resolt
		to = to / 10
		n = to 

	}
	return fi + resolt
}