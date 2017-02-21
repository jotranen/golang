package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io"
)

/*
 * Exercise 1.10: Find a web site that produces a large amount of data. Investigate caching by running fetchall
 * twice in succession to see whether the reported time changes much. Do you get the same content each time?
 * Modify fetchall to print its output to a file so it can be examined.
 */

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Printf(<-ch)
	}
	fmt.Printf(" %.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	f, err := os.Create("/tmp/dat2" + time.Now().String())

	nbytes, err := io.Copy(f, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
