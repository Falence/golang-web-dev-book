package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler logic into a closure
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to GOLANG: The Course.")
}

func main() {
	mux := http.NewServeMux()

	// Use the shortcut method ServerMux.HandleFunc
	mux.HandleFunc("/welcome", messageHandler)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}