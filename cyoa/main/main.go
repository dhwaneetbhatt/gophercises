package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/dhwaneetbhatt/gophercises/cyoa"
)

func checkAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	port := flag.Int("port", 3000, "the port to start the server on")
	filename := flag.String("JSON file", "gopher.json", "path to JSON file containing the story")
	templatePath := flag.String("template file path", "default.tmpl", "path for the HTML template file for rendering the story")
	flag.Parse()

	reader, err := os.Open(*filename)
	checkAndPanic(err)

	story, err := cyoa.JSONStory(reader)
	checkAndPanic(err)

	template, err := parseTemplate(*templatePath)
	checkAndPanic(err)

	handler := cyoa.NewHandler(story, cyoa.WithCustomTemplate(template))
	startServer(*port, story, handler)
}

func startServer(port int, story cyoa.Story, handler http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	fmt.Printf("Starting the server on port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler), mux)
}

func parseTemplate(templatePath string) (*template.Template, error) {
	bytes, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return nil, err
	}
	return template.Must(template.New("").Parse(string(bytes))), nil
}
