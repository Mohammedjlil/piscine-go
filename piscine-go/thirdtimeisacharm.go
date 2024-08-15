package piscine

func ThirdTimeIsACharm(str string) string {
/* 	if len(str) <  3 {
		return "" + "\n"
	}
	co := ""
	for i := 2 ; i < len(str) ;{
		co += string(rune(str[i]))
		i+= 3
	} 
	return co + "\n" */
	res:= ""
	for i := 2 ; i < len(str) ; i+=3{
		res +=string(str[i])
	}
	return res
}

