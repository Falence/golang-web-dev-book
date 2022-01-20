package main

import (
	// "fmt"
	// "log"
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

	// 1 ==================================
	// Embedding child collection
	// doc := Category{
	// 	bson.NewObjectId(),
	// 	"Open Source",
	// 	"Tasks for open source projects",
	// 	[]Task{
	// 		{"Create project in mgo", time.Date(2015, time.August, 10, 0, 0, 0, 0, time.UTC)},
	// 		{"Create REST API", time.Date(2015, time.August, 20, 0, 0, 0, 0, time.UTC)},
	// 	},
	// }

	// Insert a category object with embedded tasks
	// err = c.Insert(&doc)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 2 ============================================
	// Retrieving all records
	// iter := c.Find(nil).Sort("name").Iter()
	// result := Category{}
	// for iter.Next(&result) {
	// 	fmt.Printf("Category:%s, Description:%s\n", result.Name, result.Description)
	// 	tasks := result.Tasks
	// 	for _, v := range tasks {
	// 		fmt.Printf("Task:%s Due:%v\n", v.Description, v.Due)
	// 	}
	// }

	// if err = iter.Close(); err != nil {
	// 	log.Fatal(err)
	// }

	// 3 ================================================
	// Retrieving a single record from a collection
	// result := Category{}
	// err = c.Find(bson.M{"name": "Open Source"}).One(&result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Category:%s, Description:%s\n", result.Name, result.Description)
	// tasks := result.Tasks
	// for _, v := range tasks {
	// 	fmt.Printf("Tasks:%s Due:%s", v.Description, v.Due)
	// }

	// 4 =================================================
	// Updating a document
	id := "61e9b755fc51a88dd58e6d58"
	err = c.Update(bson.M{"_id": id},
	bson.M{"$set": bson.M{
		"description": "Create open-source projects",
		"tasks": []Task{
			{"Evaluate Negroni Project", time.Date(2015, time.August, 15, 0, 0, 0, 
				0, time.UTC)},
			{"Explore mgo Project", time.Date(2015, time.August, 10, 0, 0, 0, 0, 
				time.UTC)},
			{"Explore Gorilla Toolkit", time.Date(2015, time.August, 10, 0, 0, 0, 0, 
				time.UTC)},
		},
	}})
}