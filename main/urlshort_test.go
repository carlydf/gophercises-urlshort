package main

import (
	"fmt"
	"testing"
	"github.com/gophercises/urlshort"
)

func TestReadErr(t *testing.T) {
	mux := defaultMux()
	pathsToUrls := map[string]string{}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)
	badFilePath := "urls.yamls"
	_, err := urlshort.YAMLHandler(badFilePath, mapHandler)
	fmt.Println("reading bad file path...")
	fmt.Println(err)
	if err == nil {
		t.Error()
	}
}
