package main

import (
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

}
