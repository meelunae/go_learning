package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v3"
)

type URLPath struct {
	Path string `json:"path" yaml:"path"`
	URL  string `json:"url" yaml:"url`
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
	// Converting YAML byte stream to slice of structs
	yd, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	// Creating map in format wanted by our MapHandler
	pathsToUrls := buildMapFromURLStructs(yd)
	// Returning a MapHandler
	return MapHandler(pathsToUrls, fallback), nil
}

func parseYAML(data []byte) ([]URLPath, error) {
	var y []URLPath
	err := yaml.Unmarshal(data, &y)
	if err != nil {
		return nil, err
	}
	return y, nil
}

func JSONHandler(jsonData []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Converting JSON byte stream to slice of structs
	jd, err := parseJSON(jsonData)
	if err != nil {
		return nil, err
	}
	// Creating map in format wanted by our MapHandler
	pathsToUrls := buildMapFromURLStructs(jd)
	// Returning a MapHandler
	return MapHandler(pathsToUrls, fallback), nil
}

func parseJSON(data []byte) ([]URLPath, error) {
	var j []URLPath
	err := json.Unmarshal(data, &j)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func buildMapFromURLStructs(urls []URLPath) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, u := range urls {
		pathsToUrls[u.Path] = u.URL
	}
	return pathsToUrls
}
