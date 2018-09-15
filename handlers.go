package urlshortener

import (
	"net/http"
	"fmt"
	yaml "gopkg.in/yaml.v2"
)

func MapHandler(paths map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		reqPath := req.URL.Path[1:]

		fullPath, ok := paths[reqPath]
		if !ok {
			if fallback == nil {
				http.NotFound(res, req)
			} else {
				fallback.ServeHTTP(res, req)
			}
			return
		}
		http.Redirect(res, req, fullPath, http.StatusFound)
	}
}

// Yaml format is - path: test
//					url: test.com
func YamlHandler(yamlIn []byte, fallback http.Handler) http.HandlerFunc {
	m, err := parseYaml(yamlIn)
	if err != nil {
		fmt.Println(err)
		return fallback.ServeHTTP
	}
	return MapHandler(m, fallback)
}

func parseYaml(yamlIn []byte) (map[string]string, error) {
	var pathList []path
	m := make(map[string]string)
	err := yaml.Unmarshal(yamlIn, &pathList)
	if err != nil {
		return nil, err
	}

	// Create map from list
	for _, p := range(pathList) {
		m[p.Path] = p.URL
	}
	fmt.Println(m)
	return m, nil
}

type path struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}