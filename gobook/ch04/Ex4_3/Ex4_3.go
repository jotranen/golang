package main

import (
	"log"
)

/*
 * Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.
 */

func main() {
	a := [...]int{0,1,2,3,4,5}
	b := a[:]
	reverse(&b)
	log.Printf("Pointer:  %v", b)

}

func reverse(ptr *[]int) {
	for i, j := 0, len(*ptr) -1; i <j; i, j = i+1, j-1 {
		(*ptr)[i], (*ptr)[j] = (*ptr)[j], (*ptr)[i]
	}
}