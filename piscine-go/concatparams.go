package piscine

func ConcatParams(args []string) string {
	if len(args) == 0 {
		return ""
	}
	alpha := args[0]
	for i := 1; i < len(args); i++ {
		alpha += "\n" + args[i]
	}
	return alpha
}
