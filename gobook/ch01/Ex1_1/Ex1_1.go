package main

import (
	"os"
	"fmt"
)

func main() {
	s, sep := os.Args[0] + " ", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
