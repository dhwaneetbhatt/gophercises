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

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/my-github":   "https://github.com/dhwaneetbhatt",
		"/gophercises": "https://gophercises.com",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yaml, err := ioutil.ReadFile(yamlFile)
	check(err)
	yamlHandler, err := urlshort.YAMLHandler(yaml, mapHandler)
	check(err)

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
