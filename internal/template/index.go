package template

import (
	"html/template"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, req *http.Request) {
	tmpl, err := template.ParseFiles("internal/template/index.html")
	if err != nil {
		log.Printf("%+v", err)
		http.Error(w, "Couldn't parse template", http.StatusInternalServerError)
		return
	}
	log.Println("parsed template")
	err = tmpl.Execute(w, "")
	if err != nil {
		log.Printf("%+v", err)
		http.Error(w, "couldn't execute tmpl", http.StatusInternalServerError)
		return
	}
	log.Println("executed template")
}