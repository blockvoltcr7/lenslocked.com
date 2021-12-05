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
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on 3000:")
	http.ListenAndServe(":3000", r)

}
