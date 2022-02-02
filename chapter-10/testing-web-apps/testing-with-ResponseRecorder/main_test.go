package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// User Story - Users should be able to view list of User entity
func TestGetUsers(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", getUsers).Methods("GET")
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Error(err)
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("HTTP status expected: 200, got: %d", w.Code)
	}
}

// User Story - Users should be able to create a User entity
func TestCreateUser(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/users", createUser).Methods("POST")

	userJson := `{"firstname": "Falence", "lastname": "Lemungoh", "email": "falence@lemu.com"}`
	req, err := http.NewRequest(
		"POST",
		"/users",
		strings.NewReader(userJson),
	)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != 201 {
		t.Errorf("HTTP status expected: 201, got: %d", w.Code)
	}
}