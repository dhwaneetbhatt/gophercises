package link

import (
	"io"

	"golang.org/x/net/html"
)

// Link represent the link tag in the HTML document
type Link struct {
	Href string
	text string
}

// Parse returns list of parsed links from the HTML document
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, parseLink(node))
	}
	return links, nil
}

func linkNodes(node *html.Node) []*html.Node {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}
	var nodes []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, linkNodes(c)...)
	}
	return nodes
}

func parseLink(node *html.Node) Link {
	link := Link{}
	for _, a := range node.Attr {
		if a.Key == "href" {
			link.Href = a.Val
			break
		}
	}
	return link
}
