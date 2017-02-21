package main

import (
	"fmt"
	"jotranen/golang/gobook/ch02/Ex2_1/lib"
)

/*
 * Exercise 2.1: Add types, constants, and functions to tempconv for processing temperatures in the Kelvin scale,
 * where zero Kelvin is − 273.15 ° C and a difference of 1K has the same magnitude as 1 ° C.
 */

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Printf("Brrrr! %v\n", tempconv.CToK(tempconv.AbsoluteZeroC))

}
