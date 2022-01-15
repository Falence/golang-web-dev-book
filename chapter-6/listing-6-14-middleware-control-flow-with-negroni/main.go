package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
)

func middlewareFirst(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("MiddlewareFirst - Before")
	next(w, r)
	log.Println("MiddlewareFirst - After")
}

func middlewareSecond(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	log.Println("MiddlewareSecond - Before Handler")
	if r.URL.Path == "/message" {
		if r.URL.Query().Get("password") == "pass123" {
			log.Println("Authorized to the system")
			next(w, r)
		} else {
			log.Println("Failed to authorize to the system")
			return
		}
	} else {
		next(w, r)
	}
	log.Println("MiddlewareSecond - After Handler")
}

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing index Handler")
	fmt.Fprintf(w, "Welcome")
}

func message(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing message Handler")
	fmt.Fprintf(w, "HTTP middleware is awesome")
}



func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/message", message)

	n := negroni.Classic()	// creating a negroni instance
	n.Use(negroni.HandlerFunc(middlewareFirst))
	n.Use(negroni.HandlerFunc(middlewareSecond))
	n.UseHandler(mux)
	n.Run(":8080")
}