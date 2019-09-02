package cyoa

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

var htmlStoryTpl = `<!DOCTYPE html>
<html lang="en">
	<head>
	<title>Choose Your Own Adventure</title>
		<meta charset="utf-8">
		</head>
		<body>
		<h1>{{.Title}}</h1>
		{{range .Paragraphs}}
			<p>{{.}}</p>
		{{end}}
		{{range .Options}}
			<li><a href="/{{.Chapter}}">{{.Text}}</a></li>
			{{end}}
		</body>
</html>`

func init() {
	tpl = template.Must(template.New("").Parse(htmlStoryTpl))
}

// HandlerOption is a function that can modify the handler
type HandlerOption func(h *handler)

// WithCustomTemplate customizes the template to display the story
func WithCustomTemplate(t *template.Template) HandlerOption {
	return func(h *handler) {
		h.template = t
	}
}

// WithURLParserFn customizes the path parsing for the story
func WithURLParserFn(fn func(r *http.Request) string) HandlerOption {
	return func(h *handler) {
		h.urlParserFn = fn
	}
}

// NewHandler returns a http.Handler that can present the story
func NewHandler(story Story, opts ...HandlerOption) http.Handler {
	h := handler{story, tpl, defaultPathFn}
	for _, opt := range opts {
		opt(&h)
	}
	return h
}

func defaultPathFn(r *http.Request) string {
	path := r.URL.Path
	if path == "" || path == "/" {
		path = "/intro"
	}
	// removing the leading slash
	return path[1:]
}

type handler struct {
	story    Story
	template *template.Template
	urlParserFn   func(r *http.Request) string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.urlParserFn(r)
	if chapter, ok := h.story[path]; ok {
		err := h.template.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
		return
	}
	http.Error(w, "Story not found", http.StatusNotFound)
}
