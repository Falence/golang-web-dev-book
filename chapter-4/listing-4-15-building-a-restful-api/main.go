package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedOn   time.Time `json:"createdon"`
}

// Store for the Notes collection
var noteStore = make(map[string]Note)

// Variable to generate key for the collection
var id int = 0

// HTTP Post - /api/notes
func PostNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	// Decode the incoming Note json
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}

	// Store note to noteStore
	note.CreatedOn = time.Now()
	id++
	k := strconv.Itoa(id)
	noteStore[k] = note

	// Send response
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// HTTP Get - /api/notes
func GetNoteHandler(w http.ResponseWriter, r *http.Request) {
	// Get all notes
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}

	// Send response
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// HTTP Put - /api/notes/{id}
func PutNoteHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	vars := mux.Vars(r)
	k := vars["id"]
	var noteToUpd Note

	// Decode the incoming Note json
	err = json.NewDecoder(r.Body).Decode(&noteToUpd)
	if err != nil {
		panic(err)
	}

	// Check if note exists
	if note, ok := noteStore[k]; ok {
		noteToUpd.CreatedOn = note.CreatedOn
		// Delete existing item and add the updated item
		delete(noteStore, k)
		noteStore[k] = noteToUpd
	} else {
		log.Printf("Could not find key of Note %s to delete", k)
	}
	w.WriteHeader(http.StatusNoContent)
}
