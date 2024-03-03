package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

type YAMLUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

type JSONUrl struct {
	Path string `json:"path"`
	URL  string `json:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// We extract path from the request, if it matches a path contained in our map we redirect
		// otherwise, we go to fallback.
		path := r.URL.Path
		// This syntax assigns the value to dest if it exists, and the result in ok.
		// ; ok { code } is shorthand form for: if ok { code }
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
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
//   - path: /some-path
//     url: https://www.some-url.com/demo
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var y []YAMLUrl
	err := yaml.Unmarshal(yml, &y)
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Iterating through each struct unmarshaled from the YAML input to find a match
		// if we fail to find one, we go through the fallback chain.
		for _, u := range y {
			if u.Path == path {
				http.Redirect(w, r, u.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var j []JSONUrl
	err := json.Unmarshal(jsonData, &j)
	if err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// Iterating through each struct unmarshaled from the JSON input to find a match
		// if we fail to find one, we go through the fallback chain.
		for _, u := range j {
			if u.Path == path {
				http.Redirect(w, r, u.URL, http.StatusFound)
				return
			}
		}
		fallback.ServeHTTP(w, r)
	}, nil
}
