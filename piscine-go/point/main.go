package main

import "github.com/01-edu/z01"

const s = "x = 42, y = 21\n"

func main() {
	for _, v := range s {
		z01.PrintRune(v)
	}
}
