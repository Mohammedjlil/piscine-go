package main

import "fmt"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	fo := "0123456789abcdef"
	to := ""
	do := 1
	for _, v := range arr {
		if do == 4 {
			to += string(fo[v/16]) + string(fo[v%16])
			fmt.Println(to)
			to = ""
			do = 1
		} else {
			to += string(fo[v/16]) + string(fo[v%16])+ ""
			do++
		}

	}
	fmt.Println(to)
	for _,x := range arr {
		if x >= 32 && x<= 126 {
			fmt.Print(string(x))
		} else {
			fmt.Print(".")
		}


	}
	fmt.Println()
}
