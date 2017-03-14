package main

import (
	"fmt"
	"strings"
	"bufio"
	"bytes"
)

/*
 * Exercise 7.1: Using the ideas from ByteCounter, implement counters for words and for lines.
 * You will find bufio.ScanWords useful.
 */

type ByteCounter int

func (c* ByteCounter) Write(p []byte)(int, error) {
	*c += ByteCounter(len(p))

	return len(p), nil
}

type WordCounter int

func(wc* WordCounter) Write(p []byte)(int, error) {

	words := strings.Fields(string(p[:]))
	for range words {
		*wc += 1
	}
	return int(*wc), nil
}

type LineCounter int

func(lc* LineCounter) Write(p []byte)(int, error) {
	r := bytes.NewReader(p)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lc += 1
	}

	return int(*lc), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Printf("bytes: %d\n", c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Printf("bytes: %d\n", c)

	var wc WordCounter
	wc.Write([]byte("kala kissa"))
	fmt.Printf("words: %d\n", wc)

	var lc LineCounter
	lc.Write([]byte("kala kissa"))
	fmt.Printf("lines: %d\n", lc)

	lc = 0
	lc.Write([]byte("kala kissa\nkeke\nkissa"))
	fmt.Printf("lines: %d\n", lc)
}
