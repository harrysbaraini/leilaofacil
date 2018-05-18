package users

import "gopkg.in/mgo.v2/bson"

type User struct{
	ID bson.ObjectId `bson: "_id"`
	Name string `json:"name"`
	CreatedAt bson.MongoTimestamp `json:"createdAt"`
	UpdatedAt bson.MongoTimestamp `json:"updatedAt"`
}

type Users []User