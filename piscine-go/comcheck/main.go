package main

import (
	"fmt"
	"os"
)

func main() {
	arges := os.Args[1:]
	for _, arg := range arges {
		if arg == "01" || arg == "galaxy" || arg == "galaxy 01" {
			fmt.Println("Alert!!!")
			return
		}
	}
}
