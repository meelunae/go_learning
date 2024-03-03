package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"workspace/urlshort"
)

func main() {
	fp := flag.String("file", "", "file we want to load paths from")
	flag.Parse()
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	// Reading file content to byte stream that we use input for YAMLHandler
	bs, err := os.ReadFile(*fp)
	if err != nil {
		fmt.Println("Error while reading file: ", err)
	}

	// Creating separate URL Handlers depending on the input we are considering
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	yamlHandler, err := urlshort.YAMLHandler(bs, mapHandler)
	if err != nil {
		fmt.Println("Error while creating YAML handler: ", err)
	}
	jsonHandler, err := urlshort.JSONHandler(bs, yamlHandler)
	if err != nil {
		fmt.Println("Error while creating JSON handler: ", err)
	}

	fmt.Println("Starting the server on :8080")
	if jsonHandler != nil {
		fmt.Println("Running with JSON handler as primary...")
		http.ListenAndServe(":8080", jsonHandler)
	} else if yamlHandler != nil {
		fmt.Println("Running with YAML handler as primary...")
		http.ListenAndServe(":8080", yamlHandler)
	} else {
		fmt.Println("Running with map handler as primary...")
		http.ListenAndServe(":8080", mapHandler)
	}

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
