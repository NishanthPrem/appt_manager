package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("index.html"))
var dataTmpl = template.Must(template.New("data").Parse(`<p>See I told you!</p>`))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	if err := dataTmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/data", dataHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Println("Starting server on port localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
