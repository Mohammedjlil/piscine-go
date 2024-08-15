package piscine

func LastWord(s string) string {
	k := len(s)
	if k == 0 {
		return "\n"
	}
	end := k - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end == 0 {
		return "\n"
	}
	star := end
	for star >= 0 &&  s[star] != ' ' {
		star --
	} 
	return s[star+1 : end+1] + "\n"
}
