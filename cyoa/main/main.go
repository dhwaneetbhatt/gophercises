package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/dhwaneetbhatt/gophercises/cyoa"
)

func checkAndPanic(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	storyPath := flag.String("story", "gopher.json", "path to JSON file containing the story")
	templatePath := flag.String("template", "default.tmpl", "path for the HTML template file for rendering the story")
	port := flag.Int("port", 3000, "the port to start the server on")
	flag.Parse()

	// parse settings, create a new HTTP handler and start the server
	settings := cyoa.Settings{StoryFilePath: *storyPath, TemplatePath: *templatePath}
	handler, err := cyoa.NewHandler(settings)
	checkAndPanic(err)
	startServer(*port, handler)
}

func startServer(port int, handler http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/", handler)
	fmt.Printf("Starting the server on port: %d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler), mux)
}
