package models

import "gopkg.in/mgo.v2/bson"

// "encoding/json"

//MODEL OF USER -- EVERY USER IN PROJECT WILL HAVE THESE 4 ATTRIBUTES
type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}
