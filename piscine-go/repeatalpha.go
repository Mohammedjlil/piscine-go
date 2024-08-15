package piscine

func RepeatAlpha(s string) string {
	srt := ""
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			r := v - 'a' + 1
			for range r {
				srt += string (v)
			}		
		} else if  v >= 'A' && v <= 'Z' {
			r := v - 'A' + 1
			for range r {
				
			srt += string (v)
			}	
		}
	}
	return srt
}