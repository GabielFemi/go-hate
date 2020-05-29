package hateSpeech

import (
	"html/template"
	"log"
	"net/http"
)

func render (w http.ResponseWriter, tmpl string,r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(w, nil)
}