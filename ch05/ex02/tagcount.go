// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func visitTag(n *html.Node, tagCount map[string]int) {
	if n.Type == html.ElementNode {
		tagCount[n.Data]++
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitTag(c, tagCount)
	}
}

func main() {
	counts := make(map[string]int)

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "tagcount: %v\n", err)
		os.Exit(1)
	}

	visitTag(doc, counts)

	for t, c := range counts {
		fmt.Printf("%10s: %d\n", t, c)
	}
}
