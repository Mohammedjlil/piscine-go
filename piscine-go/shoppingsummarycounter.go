package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	koko := make(map[string]int)
	var ara string
	for _, e := range str {
		if e == 32 {
			koko[ara] += 1
			ara = ""
		} else if e != 32 {
			ara += string(byte(e))
		}
	}
	koko[ara] += 1

	return koko
}
