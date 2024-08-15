package piscine

func FirstWord(s string) string {
	resolt := ""
	for _, v := range s {
		if resolt != "" && v == ' ' {
			break
		} else if v != ' ' {
			resolt += string(v)
		}
	}
	return resolt + "\n"
}
