package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Description string
	Due time.Time
}

type Category struct {
	Id bson.ObjectId `bson:"_id,omitempty"`
	Name string
	Description string
	Tasks []Task
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	// Get collection
	c := session.DB("taskdb").C("categories")

	// Embedding child collection
	doc := Category{
		bson.NewObjectId(),
		"Open Source",
		"Tasks for open source projects",
		[]Task{
			{"Create project in mgo", time.Date(2015, time.August, 10, 0, 0, 0, 0, time.UTC)},
			{"Create REST API", time.Date(2015, time.August, 20, 0, 0, 0, 0, time.UTC)},
		},
	}

	// Insert a category object with embedded tasks
	err = c.Insert(&doc)
	if err != nil {
		log.Fatal(err)
	}
}