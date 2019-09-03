package cyoa

import (
	"encoding/json"
	"os"
)

// ParseStory returns a story from a JSON file
func ParseStory(storyFilePath string) (Story, error) {
	reader, err := os.Open(storyFilePath)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(reader)
	var story Story
	err = decoder.Decode(&story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

// GetChapter returns the next chapter from the path
func (story Story) GetChapter(path string) string {
	if path == "" || path == "/" {
		path = getDefaultChapter()
	}
	return path
}

func getDefaultChapter() string {
	return "intro"
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
