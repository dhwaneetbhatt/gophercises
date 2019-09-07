package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dhwaneetbhatt/gophercises/link"
)

// Link represent the a tag in the HTML document
type Link struct {
	Href string
	text string
}

func checkAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	htmlFilePath := flag.String("html", "ex1.html", "path to the HTML file")
	flag.Parse()

	r, err := os.Open(*htmlFilePath)
	checkAndPanic(err)

	links, err := link.Parse(r)
	checkAndPanic(err)
	fmt.Printf("%+v\n", links)
}
