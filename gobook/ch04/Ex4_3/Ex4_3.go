package main

import (
	"log"
)

/*
 * Exercise 4.3: Rewrite reverse to use an array pointer instead of a slice.
 */

func main() {
	a := [...]int{0,1,2,3,4,5}
	reverse(a[:])
	log.Printf("Original: %v", a)

	a = [...]int{0,1,2,3,4,5}
	b := a[:]
	reverse4_3(&b)
	log.Printf("Pointer:  %v", b)

}

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i,j  = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverse4_3(ptr *[]int) {
	for i, j := 0, len(*ptr) -1; i <j; i, j = i+1, j-1 {
		(*ptr)[i], (*ptr)[j] = (*ptr)[j], (*ptr)[i]
	}
}