package main

import (
	"os"
	"fmt"
	"strconv"
)

/*
 * Exercise 1.2: Modify the echo program to print the index and value of each of its arguments, one per line.
 */

func main() {
	s, nl := "", ""
	for index, arg := range os.Args[1:] {
		s += nl + strconv.Itoa(index) + " " + arg
		nl = "\n"
	}
	fmt.Println(s)
}
