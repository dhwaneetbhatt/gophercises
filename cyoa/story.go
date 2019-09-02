package cyoa

import (
	"encoding/json"
	"io"
)

// JSONStory returns a story from a JSON file
func JSONStory(r io.Reader) (Story, error) {
	decoder := json.NewDecoder(r)
	var story Story
	err := decoder.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

// Story which defined the mapping from chaper names to Story
type Story map[string]Chapter

// Chapter in a story
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option for next chaper
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}
