package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
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

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	type ParsedYAML struct {
		Path string
		Url string
	}
	var arr [2]ParsedYAML
	err := yaml.Unmarshal(yml, &arr)
	pathsToUrls := map[string]string{}
	for i := 0; i < len(arr); i++ {
		p := arr[i]
		pathsToUrls[p.Path] = p.Url
	}
	return func(w http.ResponseWriter, r *http.Request) {
		destination, present := pathsToUrls[r.URL.Path]
		if !present {
			fallback.ServeHTTP(w, r)
			return
		}
		http.Redirect(w, r, destination, 302)
	}, err
}
