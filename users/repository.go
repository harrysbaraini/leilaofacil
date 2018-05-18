package users

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

// Repository ...
type Repository struct{}

// SERVER determines the MongoDB Server address
const SERVER = "localhost:27017"

// DBNAME determines the database name
const DBNAME = "default"

// COLNAME determines the collection name
const COLNAME = "users"

func (r Repository) GetUsers() Users {
	session, err := mgo.Dial(SERVER)

	if err != nil {
		fmt.Println("Failed to establish connection to Mongo Server:", err)
	}

	defer session.Close()

	c := session.DB(DBNAME).C(COLNAME)
	results := Users{}

	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to get documents:", err)
	}

	return results
}