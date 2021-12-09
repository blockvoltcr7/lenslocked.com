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
	r.Get("/", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS,
		"home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(views.Must(views.ParseFS(templates.FS,
		"contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS,
		"faq.gohtml", "tailwind.gohtml"))))

	usersC := controllers.Users{}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)    //this is just getting the signup page
	r.Post("/users", usersC.Create) //this post function that is inside the signup page

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on 3000:")
	http.ListenAndServe(":3000", r)

}
