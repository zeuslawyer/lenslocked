package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact Me Here</h1>")
}

func noResource404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "<h1>Error: %s</h1>", http.StatusText(http.StatusNotFound))
	fmt.Fprintf(w, "<h1>404: No resource found at %q </h1>", r.URL.Path)
	fmt.Fprintf(w, "<h1>404: No resource found at %q </h1>", r.URL.RawPath)
}

func routeHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	switch path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		// http.NotFound(w, r)
		// http.Error(w, fmt.Sprintf("Hmmm...resource at path %q was not found.", r.URL.Path), http.StatusNotFound)
		noResource404Handler(w, r)
	}
}

func main() {
	http.HandleFunc("/", routeHandler)
	// http.HandleFunc("/contact", contactHandler)
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", nil)
}
