package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gophercises/urlshort"
)

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	filePath := "../urls.yaml"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	yamlHandlerFunc, err := urlshort.YAMLHandler(filePath, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandlerFunc)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", start)
	return mux
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go to localhost:8080/request to request a new short URL")
	//fmt.Fprintln(w, "Or go to one of the pre-existing short URLs listed below:") TODO
}
