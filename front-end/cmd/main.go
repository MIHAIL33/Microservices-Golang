package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	PORT = "8000"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, "templates/test.page.gohtml")
	})

	fmt.Println("Starting front-end on port " + PORT)
	err := http.ListenAndServe(":" + PORT, nil)
	if err != nil {
		log.Panic(err)
	}
}

func render(w http.ResponseWriter, content string) {

	templates := []string {
		content,
		"templates/header.partial.gohtml",
		"templates/footer.partial.gohtml",
		"templates/base.layout.gohtml",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}