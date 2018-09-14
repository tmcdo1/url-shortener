package urlshortener

import (
	"net/http"
	// "html/template"
	//"gopkg.in/yaml.v2"
)

func MapHandler(paths map[string]string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// get the extension from req
		reqPath := req.URL.Path[1:]

		fullPath, ok := paths[reqPath]
		if !ok {
			res.WriteHeader(http.StatusNotFound)
			res.Write([]byte("Page Not Found"))
			return
		}
		http.Redirect(res, req, fullPath, http.StatusFound)
	}
}