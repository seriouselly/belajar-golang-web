package helpers

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

func LoadTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath, err := template.ParseFiles(path.Join("views", tmpl), path.Join("views/layout", "layout.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}

	err = tmplPath.Execute(w, data)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Error is happening, keep calm", http.StatusInternalServerError)
		return
	}
}
