package main

import (
	"io"
	"os"
	"fmt"
)

/*
 * Exercise 7.2: Write a function CountingWriter with the signature below that, given an io.Writer,
 * returns a new Writer that wraps the original, and a pointer to an int64 variable that
 * at any moment contains the number of bytes written to the new Writer.
 *
 * func CountingWriter( w io.Writer) (io.Writer, *int64)
 */

func main() {
	writer, cnt := CountingWriter(os.Stdout)
	fmt.Printf("Count: %d\n", *cnt)
	fmt.Fprintf(writer, "test 123\n")
	fmt.Printf("Count: %d\n", *cnt)

}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cnt := Counter{w, 0}
	return &cnt, &cnt.cnt
}

type Counter struct {
	w 	io.Writer
	cnt	int64
}

func (c *Counter) Write(p []byte)(int, error) {
	cnt, error := c.w.Write(p)
	c.cnt += int64(cnt)
	return cnt, error
}
