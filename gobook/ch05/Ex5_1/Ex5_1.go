package main

import (
	"golang.org/x/net/html"
	"os"
	"log"
)

/*
 * Exercise 5.1: Change the findlinks program to traverse the n.FirstChild linked list using recursive
 * calls to visit instead of a loop.
 */

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		log.Printf("findlinks: %v\n", err)
		os.Exit(1)
	}

	visit(doc)
}

func visit(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				log.Printf(a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
