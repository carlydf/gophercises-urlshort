package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
	"io/ioutil"
	"os"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		destination, present := pathsToUrls[r.URL.Path]
		if !present {
			fallback.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, destination, 302)
	}
}

// YAMLHandler will return an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the map, then the
// fallback http.Handler will be called instead.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return MapHandler(pathsToUrls, fallback)
}


// YAML2Map will parse the provided YAML and then return
// a map[string]string mapping urls to their paths.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
func YAML2Map(ymlFile string) (map[string]string, error) {
	content, readErr := ioutil.ReadFile(ymlFile)
	if readErr != nil {
		return nil, readErr
	}

	type ParsedYAML struct {
		Path string
		Url string
	}
	var ymlSlice []ParsedYAML
	err := yaml.Unmarshal(content, &ymlSlice)
	pathsToUrls := map[string]string{}
	for i := 0; i < len(ymlSlice); i++ {
		p := ymlSlice[i]
		pathsToUrls[p.Path] = p.Url
	}
	return pathsToUrls, err
}

func Strings2YAML(path string, url string, ymlFile string) {
	type ParsedYAML struct {
		Path string
		Url string
	}
	var ymlSlice []ParsedYAML
	ymlSlice = append(ymlSlice, ParsedYAML{Path: path, Url: url})
	outBytes, err := yaml.Marshal(ymlSlice)
	if err != nil {
		panic(err)
	}
	writeErr := ioutil.WriteFile(ymlFile, outBytes, os.ModeAppend) // it's overwriting the whole file instead of appending
	if writeErr != nil {
		fmt.Println("something went wrong with writing to the yaml")
		panic(err)
	}
}
