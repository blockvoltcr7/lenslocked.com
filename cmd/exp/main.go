package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Bio  string
	Age  int
	Job  string
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Jon Calhoun",
		Bio:  `<script>alert("Haha, you have been h4x0r3d!");</script>`,
		Age:  123,
		Job:  "Senior Go Cloud Engineer",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
