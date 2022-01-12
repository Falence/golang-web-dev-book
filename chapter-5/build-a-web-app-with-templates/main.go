package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title string
	Description string
	CreatedOn time.Time
}

// Store for the Notes collectio
var noteStore = make(map[string]Note)

// Variable to generate keys for the collection
var id int = 0


var templates map[string]*template.Template
// Compile view templates
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	templates["index"] = template.Must(template.ParseFiles("templates/index.html", "templates/base.html"))
	templates["add"] = template.Must(template.ParseFiles("templates/add.html", "templates/base.html"))
	templates["edit"] = template.Must(template.ParseFiles("templates/edit.html", "templates/base.html"))
}


func main() {
	r := mux.NewRouter().StrictSlash(false)
	fs := http.FileServer(http.Dir("public")) // creates a handler for rendering files from the public path
	
	r.Handle("/public/", fs)
	r.HandleFunc("/", getNotes)
	r.HandleFunc("/notes/add", addNote)
	r.HandleFunc("/notes/save", saveNote)
	r.HandleFunc("/notes/edit/{id}", editNote)
	r.HandleFunc("/notes/update/{id}", updateNote)
	r.HandleFunc("/notes/delete/{id}", deleteNote)

	server := &http.Server{
		Addr: ":8080",
		Handler: r,
	}
	log.Println("Listening...")
	server.ListenAndServe()

}