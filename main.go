package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"lenslocked.com/controllers"
	"lenslocked.com/templates"
	"lenslocked.com/views"
)

func main() {

	r := chi.NewRouter()

	//parse the template
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "contact.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on 3000:")
	http.ListenAndServe(":3000", r)

}
