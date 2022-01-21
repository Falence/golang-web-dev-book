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


