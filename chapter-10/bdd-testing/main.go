package main

import (
	"net/http"

	"web-dev-with-golang-book-by-shiju/chapter-10/bdd-testing/lib"
)

func main() {
	routers := lib.SetUserRoutes()
	http.ListenAndServe(":8080", routers)
}