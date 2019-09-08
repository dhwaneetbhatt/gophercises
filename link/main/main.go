package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/dhwaneetbhatt/gophercises/link"
)

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
	defer r.Close()

	links, err := link.Parse(r)
	checkAndPanic(err)
	fmt.Printf("%+v\n", links)
}
