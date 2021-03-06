/*
 * Exercise 6.1: Implement these additional methods:
 *	func (* IntSet) Len() int // return the number of elements
 *	func (* IntSet) Remove( x int) // remove x from the set
 * 	func (* IntSet) Clear() // remove all elements from the set
 * 	func (* IntSet) Copy() *IntSet // return a copy of the set
 */

package main

import (
	"bytes"
	"fmt"
)

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))

	fmt.Printf("len: %d\n", x.Len())	// len of x

	fmt.Println(x.String())			// original x
	x.Remove(144)
	fmt.Println(x.String())			// removed 144 from x

	x.Clear()				// clear x
	fmt.Println(x.String())			// removed 144 from x

	z := y.Copy()
	fmt.Println(y.String())
	fmt.Println(z.String())			// copy of y

}

type IntSet struct {
	words	[]uint64
}

// return the number of elements
func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] &= 0 << bit
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = s.words[:0]
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var n IntSet
	n = IntSet(*s)
	return &n

}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word] & (1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word & (1 << uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64 * i + j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
