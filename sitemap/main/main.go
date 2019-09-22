package main

import (
	"flag"

	"github.com/dhwaneetbhatt/gophercises/sitemap"
)

func main() {
	website := flag.String("website", "https://github.com/dhwaneetbhatt", "website URL")
	flag.Parse()
	sitemap.Get(*website)
}
