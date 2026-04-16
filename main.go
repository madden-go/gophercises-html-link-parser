package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
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
