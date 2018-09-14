package main

import (
	"net/http"
	"fmt"
	urlshort "github.com/tmcdo1/url-shortener"
)

var (
	PORT int = 8000
)

func main() {
	// Get a mux
	// mux := http.NewServeMux()

	// Get routes and redirect route
	paths := map[string]string {
		"personal": "https://tmcdo1.github.io",
	}
	// Get
	handler := urlshort.MapHandler(paths)
	// Start up HTTP server
	fmt.Printf("Your url-shortener server is running on port %d\n", PORT)
	http.ListenAndServe(fmt.Sprintf(":%d", PORT), handler)
}
