package main

import (
	"flag"

	"github.com/dhwaneetbhatt/gophercises/sitemap"
)

func main() {
	website := flag.String("website", "https://github.com/dhwaneetbhatt", "website URL")
	maxDepth := flag.Int("depth", 100, "Depth to which to traverse in the website")
	flag.Parse()
	sitemap.Get(*website, *maxDepth)
}
