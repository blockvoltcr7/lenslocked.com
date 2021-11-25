package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>i got it!!</h1>")

}

func contactHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w,
		"<h1>contact page</h1><p>to get in touch, email me at <a href=\"mailto:smaisabir212@gmail.com\">samisabir212@gmail.com</a>")

}

func pathHandler(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

// type Router struct{}

// func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {

// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found", http.StatusNotFound)
// 	}

// }

func main() {

	http.Handler - interface with the serverHTTP method
	http.HandleFunc - a function type that accepts the same args as ServeHTTP method. also implements http.Handler

	http.Handle("/", http.Handler)
	http.HandleFunc("/", pathHandler)


	fmt.Println("Starting the server on 3000:")
	err := http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
	if err != nil {
		panic(err)
	}

}
