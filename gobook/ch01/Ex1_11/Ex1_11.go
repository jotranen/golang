package main

import (
	"time"
	"os"
	"fmt"
	"net/http"
	"io/ioutil"
	"io"
	"strings"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}

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

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("\n%.2fs %7d %s", secs, nbytes, url)
}
