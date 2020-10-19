package main

import (
	"fmt"
	"github.com/gophercises/urlshort"
	"net/http"
	"os"
	"html/template"
)

var pathsToUrls map[string]string

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls = map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	hardcodedMapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	filePath := "../urls.yaml"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	var err error
	pathsToUrls, err = urlshort.YAML2Map(filePath)
	if err != nil {
		panic(err)
	}
	yamlHandlerFunc := urlshort.YAMLHandler(pathsToUrls, hardcodedMapHandler)

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandlerFunc)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", start)
	mux.HandleFunc("/request", request)
	return mux
}

func start(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go to localhost:8080/request to request a new short URL")
	fmt.Fprintln(w, "Or go to one of the pre-existing short URLs listed below:")
	for k, v := range pathsToUrls {
		fmt.Fprintln(w, k + " >>> " + v)
	}
}

func request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("../request.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		shortCode := "/" + r.Form["short_code"][0]
		longURL := r.Form["long_url"][0]
		prevURL, prs := pathsToUrls[shortCode]
		pathsToUrls[shortCode] = longURL
		if prs  {
			fmt.Fprintf(w, shortCode + " was already registered to the URL " + prevURL + ".\n")
			fmt.Fprintf(w, shortCode + " is now registered to " + longURL + ".\n")
		} else {
			fmt.Fprintf(w, shortCode + " registered successfully.") // write data to response
		}
	}
}
