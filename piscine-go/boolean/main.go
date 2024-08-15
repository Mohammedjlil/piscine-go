package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}

	z01.PrintRune('\n')
}

func isEven(nbr int) bool { // bool
	if nbr%2 == 0 { // how to know if a number is even nbr % 2 == 0
		return true // return true
	} else {
		return false // return false

		// return nbr % 2 == 0
	}
}

func main() {
	args := os.Args[1:]
	if isEven(len(args)) { // isEven(os.Arg[1:])
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
