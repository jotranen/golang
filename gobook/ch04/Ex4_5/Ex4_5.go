package main

import (
	"fmt"
	"os"
)

/*
 * Exercise 4.5: Write an in-place function to eliminate adjacent duplicates in a [] string slice.
 */

func main() {
	fmt.Printf("Args       : %s\n", os.Args[1:])
	str := elimAdjDuplicates(os.Args[1:])
	fmt.Printf("Eliminated: %s\n", str)
}

func elimAdjDuplicates(strings []string) ([] string) {

	j := 1
	for i := 1; i < len(strings); i++ {
		if strings[i] != strings[i-1] {
			fmt.Printf("here: %d\n", i)
			strings[j] = strings[i]
			j++
		}
	}
	fmt.Printf("j: %d\n", j)

	return strings[:j]
}
