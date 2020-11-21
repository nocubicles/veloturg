package utils

import (
	"html/template"
	"log"
	"net/http"
)

func Render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseGlob("./templates/*")

	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry page cannot be served", http.StatusInternalServerError)
	}

	if err := tmpl.ExecuteTemplate(w, filename, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong...", http.StatusInternalServerError)
	}
}
