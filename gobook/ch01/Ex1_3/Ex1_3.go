package main

import (
	"os"
	"fmt"
	"time"
	"strings"
)

/*
 * Exercise 1.3: Experiment to measure the difference in running time between our potentially inefficient
 * versions and the one that uses strings.Join.
 * (Section 1.6 illustrates part of the time package,
 * and Section 11.4 shows how to write benchmark tests for systematic performance evaluation.)
 */

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	elapsed1 := time.Since(start)
	fmt.Println(s)

	fmt.Println("")

	start = time.Now()
	// building a slice takes time so string join benifits start to be visible with larger (~20 words) datasets.
	sa := []string{}
	for _, arg := range os.Args[1:] {
		sa = append(sa, arg)
	}
	fmt.Println(strings.Join(sa, ""))
	elapsed2 := time.Since(start)
	fmt.Println(s)
	fmt.Println("-------------")
	fmt.Printf("*** String join took %s\n", elapsed1)
	fmt.Printf("*** String concatenation took %s\n", elapsed2)

}
