package main

import (
	"flag"
	"os"

	"github.com/dhwaneetbhatt/gophercises/sitemap"
)

func main() {
	website := flag.String("website", "https://github.com/dhwaneetbhatt", "website URL")
	maxDepth := flag.Int("depth", 5, "Depth to which to traverse in the website")
	sitemapFile := flag.String("sitemap", "sitemap.xml", "path of the sitemap xml file")
	flag.Parse()

	writer, err := os.Create(*sitemapFile)
	defer writer.Close()
	err = sitemap.Get(*website, *maxDepth, writer)
	if err != nil {
		panic(err)
	}
}
