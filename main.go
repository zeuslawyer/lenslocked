package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Me Here</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprint(w, `
	<div>
    <h2>FAQs...</h2>
	<ol>
        <b>
            <li>This is a sample Q 1?</li>
        </b>
        Here is a sample answer 1.
    <br />
        <b>
            <li>This is a sample Q 2?</li>
        </b>
        Here is a sample answer 2.
    
    </ol>
</div>
	`)
}

// Using Chi.
func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Routes
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	// r.NotFound(http.NotFound)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, fmt.Sprintf("Hmmm...resource at path %q was not found :(", r.URL.Path), http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)

}
