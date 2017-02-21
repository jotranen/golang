package main

import (
	"os"
	"strconv"
	"fmt"
	"jotranen/golang/gobook/ch02/Ex2_1/lib"
)

/*
 * Exercise 2.2: Write a general-purpose unit-conversion program analogous to cf that reads numbers
 * from its command-line arguments or from the standard input if there are no arguments,
 * and converts each number into units like temperature in Celsius and Fahrenheit, length in feet and meters,
 * weight in pounds and kilograms, and the like.
 */

func main() {
	if len(os.Args) == 1 {
		fmt.Print("Enter value: ")
		var input string
		fmt.Scanln(&input)
		process(input)

	}
	for _, arg := range os.Args[1:] {
		process(arg)
	}
}

func process(a string) {
	t, err := strconv.ParseFloat(a, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ex2_2: % %v\n", err)
		os.Exit(1)
	}

	f := tempconv.Fahrenheit(t)
	c := tempconv.Celcius(t)
	m := tempconv.Meter(t)

	fmt.Printf("%s = %s, %s = %s, %s = %s\n",
		f, tempconv.FToC(f),
		c, tempconv.CToF(c),
	        m, tempconv.MToF(m))
}