// TODO: some testing if this actually is correct...

package main

/*
 * Exercise 4.1: Write a function that counts the number of bits that are different in two SHA256 hashes.
 * (See PopCount from Section 2.6.2.)
 */

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("Number of different bytes: %d\n", numOfDifferentBytes(&c1, &c2))
}

func numOfDifferentBytes(c1 *[32]byte, c2 *[32]byte) uint {
	res := uint(0)
	for i := range c1 {
		tmp := c1[i] ^ c2[i]
		res += PopCount2_5(tmp)
	}
	return res
}

func PopCount2_5(x uint8) uint {
	var count uint
	for count = 0; x != 0; count++ {
		x &= x -1
	}
	return count
}