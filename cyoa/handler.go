package cyoa

import (
	"log"
	"net/http"
	"text/template"
)

// NewHandler returns a http.Handler that can present the story
func NewHandler(settings Settings) (http.Handler, error) {
	story, err := ParseStory(settings.StoryFilePath)
	template, err := ParseTemplate(settings.TemplatePath)
	if err != nil {
		return nil, err
	}
	h := handler{story, template, story.GetChapter}
	return h, nil
}

type handler struct {
	story        Story
	template     *template.Template
	reslovePath func(string) string
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := h.reslovePath(r.URL.Path)
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
