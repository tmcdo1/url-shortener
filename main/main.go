package main

import (
	"net/http"
	"fmt"
	// "html/template"
	urlshort "github.com/tmcdo1/url-shortener"
)

var (
	PORT int = 8000
)

func main() {
	// Default mux used as a fallback
	mux := baseMux()

	// Get routes and redirect route
	paths := map[string]string {
		"personal": "https://tmcdo1.github.io",
	}
	y := `
- path: test
  url: https://google.com
`
	// Get
	mapHandler := urlshort.MapHandler(paths, mux)
	handler := urlshort.YamlHandler([]byte(y), mapHandler)
	// Start up HTTP server
	fmt.Printf("Your url-shortener server is running on port %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), handler)
}

func baseMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", baseHandle)
	return mux
}

func baseHandle(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		w.Write([]byte("Welcome to the URL-Shortener"))
	} else {
		http.NotFound(w, r)
	}
}
