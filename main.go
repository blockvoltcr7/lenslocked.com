package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles("templates/home.gohtml")
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "there was an error parsing the templates for homeHandler", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)

	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "there was an error executing the template at homeHandler", http.StatusInternalServerError)
		return
	}

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,
		"<h1>contact page</h1><p>to get in touch, email me at <a href=\"mailto:smaisabir212@gmail.com\">samisabir212@gmail.com</a>")

}

func faqHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>FAQ page!!!</h1>")

}

// HTTP handler accessing the url routing parameters.
func MyRequestHandler(w http.ResponseWriter, r *http.Request) {
	// fetch the url parameter `"userID"` from the request of a matching
	// routing pattern. An example routing pattern could be: /users/{userID}
	userID := chi.URLParam(r, "userID")

	// fetch `"key"` from the request context
	ctx := r.Context()
	key := ctx.Value("key").(string)

	// respond to the client
	w.Write([]byte(fmt.Sprintf("hi %v, %v", userID, key)))

}

func main() {

	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/users", MyRequestHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on 3000:")
	err := http.ListenAndServe(":3000", r)

	if err != nil {
		panic(err)
	}

}
