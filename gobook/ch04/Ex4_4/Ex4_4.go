package main

import "log"

/*
 * Exercise 4.4: Write a version of rotate that operates in a single pass.
 */

func main() {
	s := []int{0,1,2,3,4,5}
	rotateLeft(s, 2)
	log.Printf("%v", s)
}

func rotateLeft(s []int, n int) {
	reverse(s[:n])
	reverse(s[n:])
	reverse(s)
}

func reverse(s []int) {
	for i, j := 0, len(s) - 1; i < j; i,j  = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}