package piscine

func JumpOver(input string) string {
	if len(input) == 0 {
		return "\n"
	}

	var result string

	for i := 2; i < len(input); i += 3 {
		result += string(input[i])
	}

	if result == "" {
		return "\n"
	}

	return result + "\n"
}
