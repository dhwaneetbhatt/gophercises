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
	htmlFilePath := flag.String("html", "", "path to the HTML file")
	flag.Parse()
	if *htmlFilePath == "" {
		fmt.Println("path to the HTML file is required")
		os.Exit(1)
	}

	r, err := os.Open(*htmlFilePath)
	checkAndPanic(err)
	defer r.Close()

	links, err := link.Parse(r)
	checkAndPanic(err)
	for _, m := range links {
		fmt.Printf("%+v\n", m)
	}
}
