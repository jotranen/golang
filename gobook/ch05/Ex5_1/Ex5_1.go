//TODO

/*
 * Exercise 5.1: Change the findlinks program to traverse the n.FirstChild linked list using recursive
 * calls to visit instead of a loop.
 */

package main

import (
	"golang.org/x/net/html"
	"os"
	"log"
)

func main() {
	doc, err := html.Parse(os.Stdin)

	if err != nil {
		log.Printf("findlinks: %v\n", err)
		os.Exit(1)
	}

	visit(nil, doc)
	//for _, link := range visit(nil, doc) {
	//	log.Printf("%v", link)
	//}
}

func visit(links []string, n *html.Node) /* []string */ {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				log.Printf(a.Val)
				//links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		//links = visit(links, c)
		visit(nil, c)
	}

	//return links
}
