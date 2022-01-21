package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

type (
	Category struct {
		Id bson.ObjectId `bson:"_id,omitempty"`
		Name string
		Description string
	}

	DataStore struct {
		session *mgo.Session
	}
)

// Close mgo.Session
func (d *DataStore) Close() {
	d.session.Close()
}

// Return a collection from the database
func (d *DataStore) C(name string) *mgo.Collection {
	return d.session.DB("taskdb").C(name)
}

// Create a new DataStore object for each HTTP request
func NewDataStore() *DataStore {
	ds := &DataStore{
		session: session.Copy(),
	}
	return ds
}


// HANDLERS
// Insert a record
func PostCategory(w http.ResponseWriter, r *http.Request) {
	var category Category
	// Decode incoming Category json
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		panic(err)
	}
	
	ds := NewDataStore()
	defer ds.Close()

	// Getting the mgo.Collection
	c := ds.C("categories")

	// Insert record
	err = c.Insert(&category)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
}