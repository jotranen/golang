package main

import (
	"os"
	"fmt"
)

/*
 * Exercise 1.1: Modify the echo program to also print os.Args[ 0], the name of the command that invoked it.
 */

func main() {
	s, sep := os.Args[0] + " ", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
