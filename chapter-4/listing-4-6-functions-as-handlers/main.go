package main

import (
	"fmt"
	"log"
	"net/http"
)

func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Golang Web Development!")
}

func main() {
	mux := http.NewServeMux()

	// Convert the messageHandler function to a handlerFun type
	mh := http.HandlerFunc(messageHandler)
	mux.Handle("/welcome", mh)

	log.Println("Listening...")
	http.ListenAndServe(":8080", mux)
}