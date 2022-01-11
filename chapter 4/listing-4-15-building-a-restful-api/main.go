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
	Title string `json:"title"`
	Description string `json:"description"` 
	CreatedOn time.Time `json:"createdon"`
}

// Store for the Notes collection
var noteStore = make(map[string]Note)

// Variable to generate key for the collection
var id int = 0