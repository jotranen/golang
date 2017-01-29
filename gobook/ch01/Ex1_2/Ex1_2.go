package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	s, nl := "", ""
	for index, arg := range os.Args[1:] {
		s += nl + strconv.Itoa(index) + " " + arg
		nl = "\n"
	}
	fmt.Println(s)
}
