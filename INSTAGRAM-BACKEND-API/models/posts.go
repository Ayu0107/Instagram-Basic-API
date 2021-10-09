package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	UserId          bson.ObjectId `json:"userId" bson:"userId"`
	Id              bson.ObjectId `json:"id" bson:"_id"`
	Caption         string        `json:"caption" bson:"caption"`
	ImageURL        string        `json:"imageUrl" bson:"imageUrl"`
	PostedTimeStamp time.Time     `json:"postedTimestamp" bson:"postedTimeStamp"`
}
