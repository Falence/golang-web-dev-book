package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handler logic into a closure
func messageHandler(message string) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	})
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/welcome", messageHandler("Welcome to Web Development in Golang!"))
	mux.Handle("/message", messageHandler("net/http is awesome"))

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}