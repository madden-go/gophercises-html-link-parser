package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type Link struct {
	Href string
	Text string
}

func main() {
	file, err := os.Open("./tests-html/ex1.html")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	doc, err := html.Parse(file)

	if err != nil {
		log.Fatal(err)
	}

	links := extractLink(doc)

	for i, link := range links {
		fmt.Printf("Link %d: Href=%s, Text=%s\n", i, link.Href, link.Text)
	}

}

func extractLink(n *html.Node) []Link {
	var links []Link

	var traverse func(n *html.Node)
	traverse = func(node *html.Node) {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			link := Link{}

			for _, a := range node.Attr {
				if a.Key == "href" {
					link.Href = a.Val
				}
			}

			link.Text = getText(node)

			links = append(links, link)
		}

		for child := node.FirstChild; child != nil; child.NextSibling {
			traverse(child)
		}
	}

	traverse(n)
	return links
}
