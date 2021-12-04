package views

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func Parse(filepath string) (Template, error) {

	tpl, err := template.ParseFiles(filepath)

	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{htmlTpl: tpl}, nil

}

func (t Template) Execute(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "there was an error executing the template",
			http.StatusInternalServerError)
		return
	}
}
