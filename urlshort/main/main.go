package main

import (
	"flag"
	"fmt"
	"github.com/dhwaneetbhatt/gophercises/urlshort"
	"io/ioutil"
	"net/http"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// parsing input flags
	var yamlFile string
	flag.StringVar(&yamlFile, "yaml", "urls.yaml", "a yaml file that contains the URL mappings")
	flag.Parse()

	// create HttpHandler
	mux := defaultMux()

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := ioutil.ReadFile(yamlFile)
	check(err)
	yamlHandler, err := urlshort.YAMLHandler(yaml, mux)
	check(err)

	// start the server and listen for requests
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}
