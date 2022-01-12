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