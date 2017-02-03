package main

import (
	"fmt"
	"time"
)

var pc[256]byte

func init() {
	for i:= range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCount2_4(x uint64) uint {
	// see performance comparison between loops and iterators
	// http://stackoverflow.com/questions/21950244/is-there-a-way-to-iterate-over-a-range-of-integers-in-golang

	var count uint = 0

	for x >  0 {
		if ((x & 1) == 1) {
			count++
		}
		x >>= 1

	}

	return count
}

func main() {
	start := time.Now()
	fmt.Printf("PopCount: %d\n", PopCount(10))
	elapsed1 := time.Since(start)

	start = time.Now()
	fmt.Printf("PopCount: %d\n", PopCount2_4(10))
	elapsed2 := time.Since(start)

	fmt.Printf("Single expression : %s\n", elapsed1)
	fmt.Printf("Looping solution  : %s\n", elapsed2)
}

