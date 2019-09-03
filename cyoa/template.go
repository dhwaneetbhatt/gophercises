package cyoa

import (
	"io/ioutil"
	"text/template"
)

// ParseTemplate parses the HTML template from the path and returns the template object
func ParseTemplate(templatePath string) (*template.Template, error) {
	bytes, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}
	return template.Must(template.New("").Parse(string(bytes))), nil
}
