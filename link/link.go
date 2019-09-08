package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Link represent the link tag in the HTML document
type Link struct {
	Href string
	Text string
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

// linkNodes does a Depth First Search on the html document and returns all link nodes
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

// parseLink parses an HTML node to retun a Link struct
func parseLink(node *html.Node) Link {
	link := Link{}
	for _, a := range node.Attr {
		if a.Key == "href" {
			link.Href = a.Val
			break
		}
	}
	link.Text = text(node)
	return link
}

// text extracts the text from the HTML node and its children
func text(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}
	if node.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.TrimSpace(ret)
}
