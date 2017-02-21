package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strings"
)

/*
 * Exercise 1.8: Modify fetch to add the prefix http:// to each argument URL if it is missing.
 * You might want to use strings.HasPrefix.
 */

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http://") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
